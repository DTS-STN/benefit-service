package lifejourneys

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	"github.com/DTS-STN/benefit-service/models"
	"github.com/labstack/gommon/log"
)

var lifeJourneys []models.LifeJourney

// The getter for questions.
// If questions
func (q ServiceStruct) LifeJourneys() []models.LifeJourney {
	if lifeJourneys == nil || len(lifeJourneys) == 0 {
		var err error
		if lifeJourneys, err = Service.LoadLifeJourneys(); err != nil {
			log.Error(err)
		}
	}
	return lifeJourneys
}

// to make following more testable, we need to do this
// trust me, I'm a developer
var osOpen = os.Open

// Loads questions from an external source
// Returns a list of questions
func (q ServiceStruct) LoadLifeJourneys() (lifeJourneys []models.LifeJourney, err error) {
	jsonFile, err := osOpen(q.Filename)

	if err != nil {
		return
	}

	defer jsonFile.Close()

	byteValue, err := readFile(jsonFile)
	if err != nil {
		return
	}

	err = json.Unmarshal(byteValue, &lifeJourneys)

	return
}

// This functions reads and returns the data from the file opened in LoadQuestions
// Accepts a reader and returns a byte array
func readFile(reader io.Reader) ([]byte, error) {
	lines, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	return lines, err
}
