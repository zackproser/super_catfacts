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

func TestMgrListCount(t *testing.T) {
	atkManager := new(AttackManager)

	testAttack, err1 := atkManager.Add(&Attack{
		Target:    "1111111111",
		StartTime: time.Now(),
	})

	if err1 != nil || testAttack == nil {
		t.Fail()
	}

	testAttack2, err2 := atkManager.Add(&Attack{
		Target:    "111111112",
		StartTime: time.Now(),
	})

	if err2 != nil || testAttack2 == nil {
		t.Fail()
	}

	atkCount := len(atkManager.List())

	if atkCount != 2 {
		t.Fail()
	}
}

func TestMgrRemove(t *testing.T) {
	atkManager := new(AttackManager)

	testAttack, err := atkManager.Add(&Attack{
		Target:    "1111111111",
		StartTime: time.Now(),
	})

	if err != nil || testAttack == nil {
		t.Fail()
	}

	atkCount := len(atkManager.List())

	if atkCount != 1 {
		t.Fail()
	}

	success, attack := atkManager.Remove("1111111111")

	if attack == nil {
		t.Fail()
	}

	if success == false {
		t.Fail()
	}

	atkCountAfterRemoval := len(atkManager.List())

	if atkCountAfterRemoval != 0 {
		t.Fail()
	}

}
