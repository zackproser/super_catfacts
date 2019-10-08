package manager

import (
	"errors"
	"fmt"
	"super_catfacts/types"
	"time"

	"github.com/ttacon/libphonenumber"
)

type AttackManager struct {
	repository []*types.Attack
}

// Run commences the attack processing subroutine
func (a *AttackManager) Run() {
	fmt.Println("AttackManager now processing attacks")
	go func() {
		for {
			time.Sleep(3 * time.Second)
			for _, atk := range a.repository {
				fmt.Printf("ID is %v , target is %v \n", atk.ID, atk.Target)
				atk.MsgCount++
				fmt.Printf("MsgCount is %v\n", atk.MsgCount)
			}
		}
	}()
}

func validateTarget(t string) (bool, string) {
	num, err := libphonenumber.Parse(t, "US")
	if err != nil {
		return false, ""
	}
	formattedNum := libphonenumber.Format(num, libphonenumber.NATIONAL)
	return true, formattedNum
}

func (a *AttackManager) attackRunning(t string) (bool, *types.Attack) {
	for _, atk := range a.repository {
		if atk.Target == t {
			return true, atk
		}
	}
	return false, nil
}

// List dumps all current attacks
func (a *AttackManager) List() []*types.Attack {
	return a.repository
}

// Lookup attempts to fetch one attack by target
func (a *AttackManager) Lookup(t string) (bool, error) {
	valid, num := validateTarget(t)
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
func (a *AttackManager) Add(atk *types.Attack) (*types.Attack, error) {
	valid, num := validateTarget(atk.Target)
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
func (a *AttackManager) Remove(t string) (bool, *types.Attack) {
	valid, num := validateTarget(t)
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
