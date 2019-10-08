package types

import "time"

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
