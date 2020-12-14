package benefits

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/DTS-STN/benefit-service/models"
	"github.com/labstack/gommon/log"
)

var benefits []models.Benefits

// Benefits is a getter for a list of benefits
func (q *BenefitsServiceStruct) Benefits() []models.Benefits {
	if benefits == nil || len(benefits) == 0 {
		var err error
		if benefits, err = BenefitsService.LoadBenefits(); err != nil {
			log.Error(err)
		}
	}
	return benefits
}

// Benefit returns a Benefit from an ID
func (q *BenefitsServiceStruct) Benefit(id string) (models.Benefits, error) {
	for _, benefit := range q.Benefits() {
		if benefit.ID == id {
			return benefit, nil
		}
	}
	return models.Benefits{}, fmt.Errorf("Cannot find Benefit with ID: %s", id)
}

// to make following more testable, we need to do this
// trust me, I'm a developer
var osOpen = os.Open

// LoadBenefits loads data from an external source
// Returns a list of questions
func (q *BenefitsServiceStruct) LoadBenefits() (benefits []models.Benefits, err error) {
	jsonFile, err := osOpen(q.Filename)

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

// SetFilePath sets the path to the file to read data from
func (q *BenefitsServiceStruct) SetFilePath(path string) {
	q.Filename = path
}

// ClearBenefits clears the underlying benefits list
func (q *BenefitsServiceStruct) ClearBenefits() {
	benefits = nil
}
