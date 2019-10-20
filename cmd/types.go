package cmd

import "time"

// AttackManager is a client for starting, stopping and tracking attacks
type AttackManager struct {
	repository []*Attack
}

// Attack represents a currently running prank against a target phone number
type Attack struct {
	ID        int       `json:id`
	Target    string    `json:target`
	StartTime time.Time `json:starttime`
	MsgCount  int       `json:msgcount`
}

// AttackResponse is a pared down representation of an attack suitable for JSON serialization, etc
type AttackResponse struct {
	ID        int       `json:id`
	Target    string    `json:target`
	StartTime time.Time `json:starttime`
	MsgCount  int       `json:msgcount`
}

// Catfacts is a slice holding fun facts about cats
type Catfacts struct {
	Facts []string
}

// Configuration is the struct that your config.yml file gets marshalled to
type Configuration struct {
	Server ServerConfiguration
	Twilio TwilioConfiguration
}

// TwilioConfiguration contains all the Twilio account info needed to run attacks
type TwilioConfiguration struct {
	Number             string
	APIKey             string
	SID                string
	MsgIntervalSeconds int
}

// ServerConfiguration contains all the info needed to administer your Catfacts service
type ServerConfiguration struct {
	FQDN             string
	CatfactsUser     string
	CatfactsPassword string
	Port             string
	Admins           []string
}

// HealthCheckResponse is returned as a sanity check when checking your service is up
type HealthCheckResponse struct {
	Heartbeat      time.Time `json:heartbeat`
	RunningAttacks int
}

// StatusResponse encodes information about currently running attacks and their targets
type StatusResponse struct {
	AttackCount int
	Targets     []string
}
