package cmd

import (
	"fmt"
	"net/http"

	"bitbucket.org/ckvist/twilio/twiml"
	"github.com/julienschmidt/httprouter"
	"github.com/kevinburke/twilio-go"
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
	invalidErr := twilio.ValidateIncomingRequest(r.Host, Config.Twilio.APIKey, r)

	if invalidErr != nil {
		//The request is not coming from Twilio - bail out
		return
	}

	sender := r.FormValue("From")
	body := r.FormValue("Body")

	log.WithFields(logrus.Fields{
		"Sender": sender,
		"Body":   body,
	}).Debug("handleInboundSMS received message")

	if isAdmin(sender) {
		handleAdminSMSRequest(sender, body, w)
	} else {
		answer := getRandomFromSlice(10000000, responses)
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
	invalidErr := twilio.ValidateIncomingRequest(r.Host, Config.Twilio.APIKey, r)

	if invalidErr != nil {
		//The request is not coming from Twilio - bail out
		return
	}

	log.Debug("handleInboundCall received call")

	resp := twiml.NewResponse()
	resp.Action(twiml.Say{
		Voice:    twiml.TwiMan,
		Language: twiml.TwiEnglishUK,
		Text:     "Thank you for calling CatFacts! Meow Meow Meow!",
	})

	resp.Send(w)
}
