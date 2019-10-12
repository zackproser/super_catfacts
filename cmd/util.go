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

func getNextCatfact(i int) string {
	if len(catfacts) == 0 {
		panic("Error loading CatFacts")
	}
	if i <= len(catfacts) {
		return catfacts[i]
	}
	return getRandomCatfact()
}

func getRandomCatfact() string {
	return catfacts[rand.Intn(len(catfacts))]
}

func getRandomAccountResponse() string {
	return responses[rand.Intn(len(responses))]
}

func validateNumber(t string) (bool, string) {
	num, err := libphonenumber.Parse(t, "US")
	if err != nil {
		return false, ""
	}
	formattedNum := libphonenumber.Format(num, libphonenumber.NATIONAL)
	return true, formattedNum
}

func renderBasicAuth() string {
	return Config.Server.CatfactsUser + ":" + Config.Server.CatfactsPassword + "@"
}

func renderServerRoot() string {
	return "https://" + renderBasicAuth() + Config.Server.FQDN
}

func getServiceStatus() *StatusResponse {
	return attackMgr.getStatus()
}
