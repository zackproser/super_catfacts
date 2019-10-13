package cmd

import (
	"testing"
	"time"
)

func TestMgrAdd(t *testing.T) {
	atkManager := new(AttackManager)

	testAttack, err := atkManager.Add(&Attack{
		Target:    "1111111111",
		StartTime: time.Now(),
	})

	if err != nil || testAttack == nil {
		t.Fail()
	}
}
