package cmd

import (
	"testing"
)

func TestLoadJSONToSlice(t *testing.T) {

	var slice []string

	catfacts := loadJSONToSlice("../data/catfacts.json", slice)

	if len(catfacts) <= 0 {
		t.Fail()
	}
}

func TestGetRandomCatfact(t *testing.T) {

	var slice = []string{
		"This is a",
		"Test slice full",
		"Of random strings",
		"And such",
	}

	str := getRandomStringFromSlice(slice)

	if str == "" {
		t.Fail()
	}

	var found = false

	for _, s := range slice {
		if s == str {
			found = true
		}
	}

	if !found {
		t.Fail()
	}

}

func TestValidateNumber(t *testing.T) {

	long := "+1 123 456 7890"
	short := "1234567890"

	longValid, longFormatted := validateNumber(long)

	if !longValid {
		t.Fail()
	}

	if longFormatted == "" {
		t.Fail()
	}

	shortValid, shortFormatted := validateNumber(short)

	if !shortValid {
		t.Fail()
	}

	if shortFormatted == "" {
		t.Fail()
	}

	if longFormatted != shortFormatted {
		t.Fail()
	}

}
