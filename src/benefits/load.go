package benefits

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
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

// Count returns a count of all benefits
func (q *ServiceStruct) Count() int {
	benefitsList := q.GetAll("")
	return len(benefitsList)
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

/*This function checks to see if the map from the request body matches a pre-defined
pattern map. Returns true if they match, else returns false
*/
func (q *ServiceStruct) Match(input, pattern map[string]interface{}) bool {

	//Iterate over map, determine if the type of the value is a string or an array, then determine equivalence
	for key, value := range pattern {
		v := reflect.ValueOf(value)
		switch v.Kind() {
		case reflect.String:
			if input[key].(string) != value.(string) { //If value doesn't match, return false
				return false
			}
		case reflect.Slice:
			var patternData []string = pattern[key].([]string)
			var inputData string = input[key].(string)
			count := 0
			for i := 0; i < len(patternData); i++ {
				if patternData[i] == inputData { //If any of the values in the array match, increment count
					count++
				}
			}
			if count == 0 { //If nothing matches, return false
				return false
			}
		}
	}
	//if maps are determined to be equivalent, return true
	return true
}
