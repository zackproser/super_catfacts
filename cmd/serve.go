package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"super_catfacts/common"
	"super_catfacts/manager"
	"super_catfacts/types"
	"time"

	"github.com/spf13/cobra"

	"github.com/julienschmidt/httprouter"
)

var attackMgr *manager.AttackManager

func init() {
	rootCmd.AddCommand(serveCommand)
}

var serveCommand = &cobra.Command{
	Use:   "serve",
	Short: "Run a Super Catfacts service",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug(common.AppName + " listening on " + common.Configuration.Server.Port)
		initServer()
	},
}

func HealthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, common.AppName+" up and running")
}

func GetAttacks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var attackResponses []*types.AttackResponse
	atks := attackMgr.List()
	for _, atk := range atks {
		attackResponses = append(attackResponses, &types.AttackResponse{
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

	attack, err := attackMgr.Add(&types.Attack{
		Target:    target,
		StartTime: time.Now(),
	})

	if err == nil {

		atkResponse := &types.AttackResponse{
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

func StopAttack(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	err := r.ParseForm()
	if err != nil {
		log.Debug("Unable to read POST values from request")
	}

	for k, v := range r.Form {
		log.Debug("key: " + k)
		log.Debug("value: " + strings.Join(v, ""))
	}

	target := r.Form.Get("target")
	if target == "" {
		fmt.Fprintf(w, "You must supply a valid target")
		return
	}

	success, attack := attackMgr.Remove(target)

	if success {
		fmt.Fprintf(w, "Successfully stopped atttack on "+attack.Target)
	} else {
		fmt.Fprintf(w, "Error stopping attack")
	}
}

func initServer() {

	attackMgr = new(manager.AttackManager)

	attackMgr.Run()

	router := httprouter.New()

	router.GET("/healthcheck", HealthCheck)

	router.GET("/attacks", GetAttacks)

	router.POST("/attacks", CreateAttack)

	router.DELETE("/attacks", StopAttack)

	log.Fatal(http.ListenAndServe(common.Configuration.Server.Port, router))
}
