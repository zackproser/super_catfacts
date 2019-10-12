package cmd

import (
	"fmt"
	"net/http"

	"bitbucket.org/ckvist/twilio/twiml"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func handleAdminSMSRequest(sender, body string, w http.ResponseWriter) {
	// Attempt to parse a target number from the body of the message
	valid, formatted := validateNumber(body)

	if valid {

		// If the number is already being attacked, stop it
		if isUnderAttack(formatted) {
			success, attack := stopAttack(formatted)
			if success {
				fmt.Fprintf(w, "Successfully terminated attack on: %v, with %v total messages sent since %v", attack.Target, attack.MsgCount, attack.StartTime)
			}
		} else {
			// Otherwise start a new attack
			atkResponse := createAttack(formatted)
			if atkResponse != nil {
				fmt.Fprintf(w, "Successfully intitiated attack on: %v, at %v", atkResponse.Target, atkResponse.StartTime)
			} else {
				fmt.Fprintf(w, "Error initializing attack on %v", body)
			}
		}
	} else {

		fmt.Fprintf(w, "Invalid attack target: %v - please supply a valid phone number", body)
	}
}

func handleInboundSMS(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sender := r.FormValue("From")
	body := r.FormValue("Body")

	log.WithFields(logrus.Fields{
		"Sender": sender,
		"Body":   body,
	}).Debug("handleInboundSMS received message")

	if isAdmin(sender) {
		handleAdminSMSRequest(sender, body, w)
	} else {
		answer := getRandomAccountResponse()
		// Further prank a non-admin user by "upgrading" their account
		resp := twiml.NewResponse()
		resp.Action(twiml.Message{
			Body: fmt.Sprintf(answer),
			From: Config.Twilio.Number,
			To:   sender,
		})
		resp.Send(w)
	}
}

func handleInboundCall(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	log.Debug("handleInboundCall received call")

	resp := twiml.NewResponse()
	resp.Action(twiml.Say{
		Voice:    twiml.TwiMan,
		Language: twiml.TwiEnglishUK,
		Text:     "Thank you for calling CatFacts!",
	})
	resp.Action(twiml.Play{
		Url: renderServerRoot() + "/static/angryMeow.wav",
	})
	resp.Action(twiml.Say{
		Voice:    twiml.TwiMan,
		Language: twiml.TwiEnglishUK,
		Text:     "Cat Facts is the number one provider of fun facts about cats! All of our representatives are currently assisting other cat lovers. Please remain on the feline! In the meantime, please listen carefully as our menu options have recently changed.",
	})

	var gatherChildren []interface{}
	gatherChildren = append(
		gatherChildren, resp.Action(twiml.Say{
			Voice:    twiml.TwiMan,
			Language: twiml.TwiEnglishUK,
			Text:     "If you would like to receive a fun cat fact right now, press 1. If you would like to learn about how you were subscribed to CAT FACTS, please press 2",
		}),
	)
	gatherChildren = append(
		gatherChildren, resp.Action(twiml.Say{
			Voice:    twiml.TwiMan,
			Language: twiml.TwiEnglishUK,
			Text:     "If for some fur-brained reason you would like to unsubscribe from fantastic hourly cat facts, please press 3 3 3 3 4 6 7 8 9 3 1 2 6 in order right now",
		}),
	)

	resp.Gather(twiml.Gather{
		Action:      renderServerRoot() + "/phonetree",
		Method:      "POST",
		FinishOnKey: "*",
		Nested:      gatherChildren,
	})

	resp.Send(w)
}

func renderPhoneTree(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	pressed := r.FormValue("Digits")
	callingUser := r.FormValue("From")

	valid, formatted := validateNumber(callingUser)

	if !valid {
		log.WithFields(logrus.Fields{
			"Pressed":      pressed,
			"Calling User": callingUser,
		}).Debug("Invalid user phone number received by renderPhoneTree")
	}

	log.WithFields(logrus.Fields{
		"Pressed":      pressed,
		"Calling User": callingUser,
	}).Debug("renderPhoneTree received call")

	resp := twiml.NewResponse()

	switch pressed {
	case "1":

		//Send the attack message
		msg, err := client.Messages.SendMessage(
			Config.Twilio.Number,
			formatted,
			getRandomCatfact(),
			nil,
		)

		if err != nil {
			log.WithFields(logrus.Fields{
				"Error": err,
			}).Debug("Error sending Twilio SMS")
		}

		if msg != nil {
			log.WithFields(logrus.Fields{
				"Status": msg.Status,
			}).Debug("Twilio SMS sent!")
		}

		resp.Action(twiml.Say{
			Voice:    twiml.TwiMan,
			Language: twiml.TwiEnglishUK,
			Text:     "One brand spanking new Cat Fact coming right up. We're working hard to deliver your fact. Thanks for using CatFacts and please call again!",
		})
		resp.Action(twiml.Play{
			Url: renderServerRoot() + "/static/shortMeow.wav",
		})

		resp.Send(w)

	case "2":
		resp.Action(twiml.Say{
			Voice:    twiml.TwiMan,
			Language: twiml.TwiEnglishUK,
			Text:     "Please wait one moment while I pull up your account",
		})
		resp.Action(twiml.Play{
			Url: renderServerRoot() + "/static/shortMeow2.wav",
		})
		resp.Action(twiml.Say{
			Voice:    twiml.TwiMan,
			Language: twiml.TwiEnglishUK,
			Text:     "Thanks for your patience. You were subscribed to CatFacts because you love fun facts about cats. As a thank you for calling in today, we will increase the frequency of your catfacts account at no extra charge. Have a furry and fantastic day!",
		})

		resp.Send(w)
	}

}
