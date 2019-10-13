package cmd

import (
	"errors"
	"math/rand"
	"time"

	"github.com/kevinburke/twilio-go"
	"github.com/sirupsen/logrus"
)

var client *twilio.Client

// Initialize loads C:atfacts
func (a *AttackManager) Initialize() {

	rand.Seed(time.Now().Unix())

	client = twilio.NewClient(Config.Twilio.SID, Config.Twilio.APIKey, nil)

	// Load CatFacts for use in attacks
	catfacts = loadJSONToSlice("data/catfacts.json", catfacts)

	// Load account responses for answering inbound SMS
	responses = loadJSONToSlice("data/responses.json", responses)
}

// Run commences the attack processing subroutine
func (a *AttackManager) Run() {

	log.WithFields(logrus.Fields{
		"Messaging Interval in Seconds": Config.Twilio.MsgIntervalSeconds,
	}).Debug(AppName + " now processing attacks")

	go func() {
		for {
			time.Sleep(time.Duration(Config.Twilio.MsgIntervalSeconds) * time.Second)
			for _, atk := range a.repository {

				//Send the attack message
				msg, err := client.Messages.SendMessage(
					Config.Twilio.Number,
					atk.Target,
					getNextCatfact(atk.MsgCount),
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
					atk.MsgCount++

				}

				log.WithFields(logrus.Fields{
					"ID":            atk.ID,
					"Target":        atk.Target,
					"Message Count": atk.MsgCount,
				}).Debug("Attack pulse")
			}
		}
	}()
}

func (a *AttackManager) getStatus() *StatusResponse {

	var targets []string

	for _, atk := range a.repository {
		targets = append(targets, atk.Target)
	}

	return &StatusResponse{
		AttackCount: len(a.repository),
		Targets:     targets,
	}
}

func (a *AttackManager) attackRunning(t string) (bool, *Attack) {
	for _, atk := range a.repository {
		if atk.Target == t {
			return true, atk
		}
	}
	return false, nil
}

// List dumps all current attacks
func (a *AttackManager) List() []*Attack {
	return a.repository
}

// Lookup attempts to fetch one attack by target
func (a *AttackManager) Lookup(t string) (bool, error) {
	valid, num := validateNumber(t)
	if valid == false {
		return false, errors.New("Invalid attack target: " + t)
	}
	for _, attack := range a.repository {
		if attack.Target == num {
			return true, nil
		}
	}
	return false, nil
}

// Add commences a new attack
func (a *AttackManager) Add(atk *Attack) (*Attack, error) {
	valid, num := validateNumber(atk.Target)
	if valid == false {
		return nil, errors.New("Invalid attack target:" + atk.Target)
	}
	running, attack := a.attackRunning(num)
	if running == true {
		return nil, errors.New("Attack already running on " + attack.Target + " count: ")
	}
	atk.Target = num
	atk.ID = len(a.repository)
	a.repository = append(a.repository, atk)
	return atk, nil
}

// Remove terminates an existing attack
func (a *AttackManager) Remove(t string) (bool, *Attack) {
	valid, num := validateNumber(t)
	if valid == false {
		return false, nil
	}
	for i, atk := range a.repository {
		if atk.Target == num {
			atk.Ticker.Stop()
			a.repository[len(a.repository)-1], a.repository[i] = a.repository[i], a.repository[len(a.repository)-1]
			a.repository = a.repository[:len(a.repository)-1]
		}
		return true, atk
	}
	return false, nil
}

// RemoveByID stops an in progress attack
func (a *AttackManager) RemoveByID(id int) (bool, *Attack) {
	for i, atk := range a.repository {
		if atk.ID == id {
			atk.Ticker.Stop()
			a.repository[len(a.repository)-1], a.repository[i] = a.repository[i], a.repository[len(a.repository)-1]
			a.repository = a.repository[:len(a.repository)-1]
			return true, atk
		}
	}
	return false, nil
}

// GetCurrentRunningAttackCount returns the number of attacks in progress
func (a *AttackManager) GetCurrentRunningAttackCount() int {
	return len(a.repository)
}
