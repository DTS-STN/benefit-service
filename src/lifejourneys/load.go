package lifejourneys

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/DTS-STN/benefit-service/models"
	"github.com/labstack/gommon/log"
)

var lifeJourneys []models.LifeJourney

// LifeJourneys returns all Life Journeys
func (q *LifeJourneyServiceStruct) LifeJourneys() []models.LifeJourney {
	if lifeJourneys == nil || len(lifeJourneys) == 0 {
		var err error
		if lifeJourneys, err = q.LoadLifeJourneys(); err != nil {
			log.Error(err)
		}
	}
	return lifeJourneys
}

// LifeJourney returns a Life Journey from an ID
func (q *LifeJourneyServiceStruct) LifeJourney(id string) (models.LifeJourney, error) {
	for _, lifeJourney := range q.LifeJourneys() {
		if lifeJourney.ID == id {
			return lifeJourney, nil
		}
	}
	return models.LifeJourney{}, fmt.Errorf("Cannot find Life Journey with ID: %s", id)
}

// to make following more testable, we need to do this
// trust me, I'm a developer
var osOpen = os.Open

// LoadLifeJourneys will get Life Journeys from an external source
// returns a list of Life Journeys
func (q *LifeJourneyServiceStruct) LoadLifeJourneys() (lifeJourneys []models.LifeJourney, err error) {
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

// readFile reads and returns the data from the file opened in LoadQuestions
// Accepts a reader and returns a byte array
func readFile(reader io.Reader) ([]byte, error) {
	lines, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	return lines, err
}

// SetFilePath sets the path to the file to read data from
func (q *LifeJourneyServiceStruct) SetFilePath(path string) {
	q.Filename = path
}

// ClearLifeJourneys clears the underlying list
func (q *LifeJourneyServiceStruct) ClearLifeJourneys() {
	lifeJourneys = nil
}
