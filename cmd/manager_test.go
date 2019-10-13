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

func TestMgrStatus(t *testing.T) {
	atkManager := new(AttackManager)

	testAttack, _ := atkManager.Add(&Attack{
		Target:    "1111111111",
		StartTime: time.Now(),
	})

	status := atkManager.getStatus()

	if status == nil {
		t.Fail()
	}

	var found = false

	for _, target := range status.Targets {
		if testAttack.Target == target {
			found = true
		}
	}

	if found == false {
		t.Fail()
	}

}

func TestIsAttackRunning(t *testing.T) {
	atkManager := new(AttackManager)

	testAttack, _ := atkManager.Add(&Attack{
		Target:    "3333333333",
		StartTime: time.Now(),
	})

	running, atk := atkManager.attackRunning(testAttack.Target)

	if !running {
		t.Fail()
	}

	if atk == nil {
		t.Fail()
	}

	testNonAddedAttack := &Attack{
		Target:    "4444444444",
		StartTime: time.Now(),
	}

	nonIsRunning, nonAtk := atkManager.attackRunning(testNonAddedAttack.Target)

	if nonIsRunning == true {
		t.Fail()
	}

	if nonAtk != nil {
		t.Fail()
	}
}
