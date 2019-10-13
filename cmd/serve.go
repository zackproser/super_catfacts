package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/julienschmidt/httprouter"
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

// HealthCheck returns a simple ping-like response
func HealthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	hc := &HealthCheckResponse{
		Heartbeat:      time.Now(),
		RunningAttacks: attackMgr.GetCurrentRunningAttackCount(),
	}
	j, err := json.Marshal(hc)
	if err != nil {
		log.WithFields(logrus.Fields{
			"Error": err,
		}).Debug("Error serializing healthcheck json")

	}
	w.Write(j)
}

// GetAttacks returns a JSON representation of all currently running attacks
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
		log.WithFields(logrus.Fields{
			"Error": err,
		}).Debug("Error serializing attack info")
	}
	w.Write(j)
}

func createAttackAPI(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

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
			log.WithFields(logrus.Fields{
				"Error": err,
			}).Debug("Error serializing attack info to JSON")
		}

		w.Write(j)
	} else {
		fmt.Fprintf(w, "Error commencing attack on ")
	}
}

func createAttack(target string) *AttackResponse {
	attack, err := attackMgr.Add(&Attack{
		Target:    target,
		StartTime: time.Now(),
	})

	if err != nil {
		log.WithFields(logrus.Fields{
			"Error": err,
		}).Debug("Error creating SMS initiated attack")
	}

	return &AttackResponse{
		ID:        attack.ID,
		Target:    attack.Target,
		StartTime: attack.StartTime,
	}
}

func stopAttack(target string) (bool, *Attack) {
	success, attack := attackMgr.Remove(target)
	if success {
		return true, attack
	} else {
		return false, nil
	}
}

func StopAttack(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	attackID := ps.ByName("id")

	if attackID == "" {
		fmt.Fprintf(w, "Must provide a valid ID of an existing attack")
	}

	attackIDInt, err := strconv.Atoi(attackID)

	if err != nil {
		log.WithFields(logrus.Fields{
			"Error": err,
		}).Debug("Error converting attack ID string to int")
	}

	success, attack := attackMgr.RemoveByID(attackIDInt)

	if success {
		fmt.Fprintf(w, "Successfully stopped atttack on "+attack.Target)
	} else {
		fmt.Fprintf(w, "Error stopping attack")
	}
}

func isAdmin(s string) bool {
	valid, formatted := validateNumber(s)
	if valid {
		for _, admin := range Config.Server.Admins {
			if formatted == admin {
				return true
			}
		}
	}
	return false
}

func isUnderAttack(s string) bool {
	valid, formatted := validateNumber(s)
	if valid {
		for _, atk := range attackMgr.repository {
			if formatted == atk.Target {
				return true
			}
		}
	}
	return false
}

func basicAuth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := r.BasicAuth()

		if hasAuth && user == Config.Server.CatfactsUser && password == Config.Server.CatfactsPassword {
			h(w, r, ps)
		} else {
			// Request Basic Authentication
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

func initServer() {

	attackMgr = new(AttackManager)

	attackMgr.Initialize()

	attackMgr.Run()

	router := httprouter.New()

	router.GET("/", basicAuth(HealthCheck))

	router.GET("/attacks", basicAuth(GetAttacks))

	router.POST("/attacks", basicAuth(createAttackAPI))

	router.DELETE("/attacks/:id", basicAuth(StopAttack))

	router.POST("/sms/receive", basicAuth(handleInboundSMS))

	router.GET("/call/receive", basicAuth(handleInboundCall))

	router.POST("/phonetree", basicAuth(renderPhoneTree))

	router.ServeFiles("/static/*filepath", http.Dir("static"))

	// Gate Catfacts API to requests that supply the correct API key in the HTTP Authorization header
	log.Fatal(http.ListenAndServe(Config.Server.Port, router))
}
