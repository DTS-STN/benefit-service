package benefits

import (
	"encoding/json"
	"github.com/DTS-STN/benefit-service/models"
	"github.com/labstack/gommon/log"
	"io"
	"io/ioutil"
	"os"
)

var benefits []models.Benefits

// The getter for questions.
// If questions
func (q BenefitsServiceStruct) Benefits() []models.Benefits {
	if benefits == nil || len(benefits) == 0 {
		var err error
		if benefits, err = BenefitsService.LoadBenefits(); err != nil {
			log.Error(err)
		}
	}
	return benefits
}

// to make following more testable, we need to do this
// trust me, I'm a developer
var osOpen = os.Open

// Loads questions from an external source
// Returns a list of questions
func (q BenefitsServiceStruct) LoadBenefits() (benefits []models.Benefits, err error) {
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
