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

var lifeJourneysEN []models.LifeJourney
var lifeJourneysFR []models.LifeJourney

var files = map[string]string{
	"en": "life_journeys_en.json",
	"fr": "life_journeys_fr.json",
}

// LifeJourneys returns all Life Journeys
func (q *LifeJourneyServiceStruct) LifeJourneys(lang string) []models.LifeJourney {
	if lang == "fr" {
		if lifeJourneysFR == nil || len(lifeJourneysFR) == 0 {
			var err error
			if lifeJourneysFR, err = q.LoadLifeJourneys(lang); err != nil {
				log.Error(err)
			}
		}
		return lifeJourneysFR
	}

	// default to english if no language is specified
	if lifeJourneysEN == nil || len(lifeJourneysEN) == 0 {
		var err error
		if lifeJourneysEN, err = q.LoadLifeJourneys(lang); err != nil {
			log.Error(err)
		}
	}
	return lifeJourneysEN
}

// LifeJourney returns a Life Journey from an ID
func (q *LifeJourneyServiceStruct) LifeJourney(lang, id string) (models.LifeJourney, error) {
	for _, lifeJourney := range q.LifeJourneys(lang) {
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
func (q *LifeJourneyServiceStruct) LoadLifeJourneys(lang string) (lifeJourneys []models.LifeJourney, err error) {
	file := files[lang]
	if file == "" {
		file = files["en"]
	}
	jsonFile, err := osOpen(file)

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
