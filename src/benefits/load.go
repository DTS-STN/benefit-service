package benefits

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/DTS-STN/benefit-service/models"
	"github.com/labstack/gommon/log"
)

var benefitsEN []models.Benefits
var benefitsFR []models.Benefits

var Files = map[string]string{
	"en": "benefit_info_en.json",
	"fr": "benefit_info_fr.json",
}

// GetAll is a getter for a list of benefits
func (q *ServiceStruct) GetAll(lang string) []models.Benefits {
	if lang == "fr" {
		if benefitsFR == nil || len(benefitsFR) == 0 {
			var err error
			if benefitsFR, err = q.LoadBenefits(lang); err != nil {
				log.Error(err)
			}
		}
		return benefitsFR
	}

	// default to english if no lang is specified
	if benefitsEN == nil || len(benefitsEN) == 0 {
		var err error
		if benefitsEN, err = q.LoadBenefits(lang); err != nil {
			log.Error(err)
		}
	}
	return benefitsEN
}

// GetByID returns a Benefit from an ID
func (q *ServiceStruct) GetByID(lang, benefitId string) (models.Benefits, error) {
	for _, benefit := range q.GetAll(lang) {
		val, err := strconv.Atoi(benefitId)
		if err != nil {
			return models.Benefits{}, fmt.Errorf("Cannot parse %s", benefitId)
		}
		if benefit.ID == val {
			return benefit, nil
		}
	}
	return models.Benefits{}, fmt.Errorf("Cannot find Benefit with ID: %s", benefitId)
}

// to make following more testable, we need to do this
// trust me, I'm a developer
var osOpen = os.Open

// LoadBenefits loads data from an external source
// Returns a list of questions
func (q *ServiceStruct) LoadBenefits(lang string) (benefits []models.Benefits, err error) {
	file := Files[lang]
	if file == "" {
		file = Files["en"]
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

	err = json.Unmarshal(byteValue, &benefits)

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
