package cmd

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/ttacon/libphonenumber"
)

func loadJSONToSlice(filePath string, s []string) []string {

	jsonFile, err := os.Open("data/catfacts.json")

	defer jsonFile.Close()

	if err != nil {
		log.WithFields(logrus.Fields{
			"Error": err,
		}).Debug("Unable to parse Catfacts JSON file")
	}

	bytes, readErr := ioutil.ReadAll(jsonFile)

	if readErr != nil {
		log.WithFields(logrus.Fields{
			"Error": err,
			"File":  filePath,
		}).Debug("Unable to read JSON file")
	}

	json.Unmarshal(bytes, &s)

	return s
}

func getRandomFromSlice(i int, s []string) string {
	if i > 0 && i < len(s) {
		return s[i]
	}
	return s[rand.Intn(len(s))]
}

func validateNumber(t string) (bool, string) {
	num, err := libphonenumber.Parse(t, "US")
	if err != nil {
		return false, ""
	}
	formattedNum := libphonenumber.Format(num, libphonenumber.NATIONAL)
	return true, formattedNum
}
