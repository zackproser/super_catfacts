// Copyright (C) 2014 Cristoffer Kvist. All rights reserved.
// This project is licensed under the terms of the MIT license in LICENSE.

// Package twiml provides Twilio Markup Language support for building web
// services with instructions for twilio how to handle incoming call or message.
package twiml

import (
	"encoding/xml"
	"fmt"
	"io"
)

// Create new response
func NewResponse() *Response {
	return new(Response)
}

// Action appends action verb structs to response. Valid verbs: Enqueue, Say,
// Leave, Message, Pause, Play, Record, Redirect, Reject, Hangup
func (r *Response) Action(structs ...interface{}) error {
	for _, s := range structs {
		switch s := s.(type) {
		default:
			return fmt.Errorf("non valid verb: '%T'", s)
		case Enqueue, Hangup, Leave, Message, Pause, Play, Record,
			Redirect, Reject, Say:
			r.Response = append(r.Response, s)
		}
	}
	return nil
}

// Dial appends dial action verb and noun structs to respose
// Valid verb: Dial. Valid nouns: Client, Conference, Number, Queue, Sip
func (r *Response) Dial(structs ...interface{}) error {
	d := Dial{}

	for _, s := range structs {
		switch s := s.(type) {
		default:
			return fmt.Errorf("non valid verb: '%T'", s)
		case Dial:
			d.HangupOnStar = s.HangupOnStar
			d.TimeLimit = s.TimeLimit
			d.CallerId = s.CallerId
			d.Timeout = s.Timeout
			d.Action = s.Action
			d.Method = s.Method
			d.Record = s.Record
			d.Number = s.Number
		case Client, Conference, Number, Queue, Sip:
			d.Nested = append(d.Nested, s)
		}
	}
	r.Response = append(r.Response, d)
	return nil
}

// Gather collects digits a caller enter by pressing the keypad
// Valid verb: Gather. Valid nested verbs: Say, Pause, Play
func (r *Response) Gather(structs ...interface{}) error {
	g := Gather{}

	for _, s := range structs {
		switch s := s.(type) {
		default:
			return fmt.Errorf("non valid verb: '%T'", s)
		case Gather:
			g.FinishOnKey = s.FinishOnKey
			g.NumDigits = s.NumDigits
			g.Timeout = s.Timeout
			g.Action = s.Action
			g.Method = s.Method
		case Say, Pause, Play: // Valid nested verbs
			g.Nested = append(g.Nested, s)
		}

	}
	r.Response = append(r.Response, g)
	return nil
}

// Send sends xml encoded response to writer
func (r Response) Send(w io.Writer) (err error) {
	enc := xml.NewEncoder(w)
	enc.Indent("  ", "   ")

	fmt.Fprintf(w, "%s", xml.Header)
	if err := enc.Encode(r); err != nil {
		return err
	}
	fmt.Fprintf(w, "\n")
	return err
}

// String returns a formatted xml response
func (r Response) String() string {
	output, err := xml.MarshalIndent(r, "  ", "    ")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	return xml.Header + string(output) + "\n"
}
