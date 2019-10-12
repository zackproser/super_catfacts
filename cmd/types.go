package cmd

import "time"

type AttackManager struct {
	repository []*Attack
}

type Attack struct {
	ID        int       `json:id`
	Target    string    `json:target`
	StartTime time.Time `json:starttime`
	MsgCount  int       `json:msgcount`
	Ticker    time.Ticker
}

type AttackResponse struct {
	ID        int       `json:id`
	Target    string    `json:target`
	StartTime time.Time `json:starttime`
	MsgCount  int       `json:msgcount`
}

type Catfacts struct {
	Facts []string
}

type Configuration struct {
	Server ServerConfiguration
	Twilio TwilioConfiguration
}

type TwilioConfiguration struct {
	Number             string
	APIKey             string
	SID                string
	MsgIntervalSeconds int
}

type ServerConfiguration struct {
	FQDN             string
	CatfactsUser     string
	CatfactsPassword string
	Port             string
	Admins           []string
}

type HealthCheckResponse struct {
	Heartbeat      time.Time `json:heartbeat`
	RunningAttacks int
}

type StatusResponse struct {
	AttackCount int
	Targets     []string
}
