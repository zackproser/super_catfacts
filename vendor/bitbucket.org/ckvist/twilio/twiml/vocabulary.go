package twiml

import "encoding/xml"

type Client struct {
	XMLName xml.Name `xml:"Client"`
	Method  string   `xml:"method,attr,omitempty"`
	Url     string   `xml:"URL,omitempty"`
	Name    string   `xml:",chardata"`
}

type Conference struct {
	XMLName                xml.Name `xml:"Conference"`
	Muted                  bool     `xml:"muted,attr,omitempty"`
	Beep                   string   `xml:"beep,attr,omitempty"`
	StartConferenceOnEnter bool     `xml:"startConferenceOnEnter,attr,omitempty"`
	EndConferenceOnExit    bool     `xml:"endConferenceOnExit,attr,omitempty"`
	WaitUrl                string   `xml:"waitUrl,attr,omitempty"`
	WaitMethod             string   `xml:"waitMethod,attr,omitempty"`
	MaxParticipants        int      `xml:"maxParticipants,attr,omitempty"`
	Name                   string   `xml:",chardata"`
}

type Dial struct {
	XMLName      xml.Name `xml:"Dial"`
	Action       string   `xml:"action,attr,omitempty"`
	Method       string   `xml:"method,attr,omitempty"`
	Timeout      int      `xml:"timeout,attr,omitempty"`
	HangupOnStar bool     `xml:"hangupOnStar,attr,omitempty"`
	TimeLimit    int      `xml:"timeLimit,attr,omitempty"`
	CallerId     string   `xml:"callerId,attr,omitempty"`
	Record       bool     `xml:"record,attr,omitempty"`
	Number       string   `xml:",chardata"`
	Nested       []interface{}
}

type Enqueue struct {
	XMLName       xml.Name `xml:"Enqueue"`
	Action        string   `xml:"action,attr,omitempty"`
	Method        string   `xml:"method,attr,omitempty"`
	WaitUrl       string   `xml:"waiUrl,attr,omitempty"`
	WaitUrlMethod string   `xml:"waiUrlMethod,attr,omitempty"`
	Name          string   `xml:",chardata"`
}

type Hangup struct {
	XMLName xml.Name `xml:"Hangup"`
}

type Leave struct {
	XMLName xml.Name `xml:"Leave"`
}

type Message struct {
	XMLName        xml.Name `xml:"Message"`
	To             string   `xml:"to,attr,omitempty"`
	From           string   `xml:"from,attr,omitempty"`
	Action         string   `xml:"action,attr,omitempty"`
	Method         string   `xml:"method,attr,omitempty"`
	StatusCallback string   `xml:"statusCallback,attr,omitempty"`
	Body           string   `xml:"Body,omitempty"`
	Media          string   `xml:"Media,omitempty"`
}

type Number struct {
	XMLName    xml.Name `xml:"Number"`
	SendDigits string   `xml:"sendDigits,attr,omitempty"`
	Url        string   `xml:"url,attr,omitempty"`
	Method     string   `xml:"method,attr,omitempty"`
	Number     string   `xml:",chardata"`
}

type Pause struct {
	XMLName xml.Name `xml:"Pause"`
	Length  int      `xml:"length,attr,omitempty"`
}

type Play struct {
	XMLName xml.Name `xml:"Play"`
	Loop    int      `xml:"loop,attr,omitempty"`
	Digits  int      `xml:"digits,attr,omitempty"`
	Url     string   `xml:",chardata"`
}

type Queue struct {
	XMLName xml.Name `xml:"Queue"`
	Url     string   `xml:"url,attr,omitempty"`
	Method  string   `xml:"method,attr,omitempty"`
	Name    string   `xml:",chardata"`
}

type Record struct {
	XMLName            xml.Name `xml:"Record"`
	Action             string   `xml:"action,attr,omitempty"`
	Method             string   `xml:"method,attr,omitempty"`
	Timeout            int      `xml:"timeout,attr,omitempty"`
	FinishOnKey        string   `xml:"finishOnKey,attr,omitempty"`
	MaxLength          int      `xml:"maxLength,attr,omitempty"`
	Transcribe         bool     `xml:"transcribe,attr,omitempty"`
	TranscribeCallback string   `xml:"transcribeCallback,attr,omitempty"`
	PlayBeep           bool     `xml:"playBeep,attr,omitempty"`
}

type Redirect struct {
	XMLName xml.Name `xml:"Redirect"`
	Method  string   `xml:"method,attr,omitempty"`
	Url     string   `xml:",chardata"`
}

type Reject struct {
	XMLName xml.Name `xml:"Reject"`
	Reason  string   `xml:"reason,attr,omitempty"`
}

type Response struct {
	XMLName  xml.Name `xml:"Response"`
	Response []interface{}
}

type Say struct {
	XMLName  xml.Name `xml:"Say"`
	Voice    string   `xml:"voice,attr,omitempty"`
	Language string   `xml:"language,attr,omitempty"`
	Loop     int      `xml:"loop,attr,omitempty"`
	Text     string   `xml:",chardata"`
}

type Sip struct {
	XMLName  xml.Name `xml:"Sip"`
	Username string   `xml:"username,attr,omitempty"`
	Password string   `xml:"password,attr,omitempty"`
	Url      string   `xml:"url,attr,omitempty"`
	Method   string   `xml:"method,attr,omitempty"`
	Address  string   `xml:",chardata"`
}

type Gather struct {
	XMLName     xml.Name `xml:"Gather"`
	Action      string   `xml:"action,attr,omitempty"`
	Method      string   `xml:"method,attr,omitempty"`
	Timeout     int      `xml:"timeout,attr,omitempty"`
	FinishOnKey string   `xml:"finishOnKey,attr,omitempty"`
	NumDigits   int      `xml:"numDigits,attr,omitempty"`
	Nested      []interface{}
}
