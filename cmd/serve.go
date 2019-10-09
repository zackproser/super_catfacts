package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/spf13/cobra"

	"github.com/julienschmidt/httprouter"

	"bitbucket.org/ckvist/twilio/twiml"
)

var attackMgr *AttackManager

func init() {
	rootCmd.AddCommand(serveCommand)
}

var serveCommand = &cobra.Command{
	Use:   "serve",
	Short: "Run a Super Catfacts service",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug(AppName + " listening on " + Config.Server.Port)
		initServer()
	},
}

func HealthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, AppName+" up and running")
}

func GetAttacks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var attackResponses []*AttackResponse
	atks := attackMgr.List()
	for _, atk := range atks {
		attackResponses = append(attackResponses, &AttackResponse{
			ID:        atk.ID,
			Target:    atk.Target,
			StartTime: atk.StartTime,
			MsgCount:  atk.MsgCount,
		})
	}
	j, err := json.Marshal(attackResponses)
	if err != nil {
		log.Debug("Error serializing attack info: %v", err)
	}
	w.Write(j)
}

func CreateAttack(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	err := r.ParseForm()
	if err != nil {
		log.Debug("Unable to read POST values from request")
	}

	target := r.Form.Get("target")
	if target == "" {
		fmt.Fprintf(w, "You must supply a valid target")
		return
	}

	attack, err := attackMgr.Add(&Attack{
		Target:    target,
		StartTime: time.Now(),
	})

	if err == nil {

		atkResponse := &AttackResponse{
			ID:        attack.ID,
			Target:    attack.Target,
			StartTime: attack.StartTime,
		}

		j, marshalErr := json.Marshal(atkResponse)

		if marshalErr != nil {
			log.Debug("Error serializing attack info to JSON: %v", marshalErr)
		}

		w.Write(j)
	} else {
		fmt.Fprintf(w, "Error commencing attack on ")
	}
}

func StopAttack(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	attackID := ps.ByName("id")

	if attackID == "" {
		fmt.Fprintf(w, "Must provide a valid ID of an existing attack")
	}

	attackIDInt, err := strconv.Atoi(attackID)

	if err != nil {
		log.Debug("Error converting attack ID string to int: %v", err)
	}

	success, attack := attackMgr.RemoveByID(attackIDInt)

	if success {
		fmt.Fprintf(w, "Successfully stopped atttack on "+attack.Target)
	} else {
		fmt.Fprintf(w, "Error stopping attack")
	}
}

func handleInboundSMS(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sender := r.FormValue("From")

	resp := twiml.NewResponse()
	resp.Action(twiml.Message{
		Body: fmt.Sprintf("Thanks for your feedback! We've awarded you additional CatFacts at no extra charge!"),
		From: Config.Twilio.Number,
		To:   sender,
	})
	resp.Send(w)
}

func initServer() {

	attackMgr = new(AttackManager)

	attackMgr.Initialize()

	attackMgr.Run()

	router := httprouter.New()

	router.GET("/healthcheck", HealthCheck)

	router.GET("/attacks", GetAttacks)

	router.POST("/attacks", CreateAttack)

	router.DELETE("/attacks/:id", StopAttack)

	router.POST("/sms/receive", handleInboundSMS)

	log.Fatal(http.ListenAndServe(Config.Server.Port, router))
}
