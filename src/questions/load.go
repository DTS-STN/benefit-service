package questions

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

var questionList = map[string][]models.Question{
	"en": {},
	"fr": {},
}

// GetAll returns a list of questions
func (q *ServiceStruct) GetAll(lang string) ([]models.Question, error) {
	// if lang isn't set, just return the english questions
	if lang == "" {
		lang = "en"
	}

	if questionList[lang] == nil || len(questionList[lang]) == 0 {
		var q []models.Question
		var err error
		if q, err = loadQuestions("questions_" + lang + ".json"); err != nil {
			return questionList[lang], err
		}
		questionList[lang] = q
	}

	return questionList[lang], nil
}

// GetByID returns a single question from an id
func (q *ServiceStruct) GetByID(lang, id string) (models.Question, error) {
	if ques, err := q.GetAll(lang); err == nil {
		for _, question := range ques {
			val, err := strconv.Atoi(id)
			if err != nil {
				return models.Question{}, fmt.Errorf("Cannot parse id: %s", id)
			}
			if question.ID == val {
				return question, nil
			}
		}
	} else {
		return models.Question{}, err
	}

	return models.Question{}, fmt.Errorf("Cannot find question with id: %s", id)
}

// to make following more testable, we need to do this
var osOpen = os.Open
var loadQuestions = loadQuestionsFunc

// LoadQuestions loads questions from an external source
// Returns a list of questions
func loadQuestionsFunc(path string) (questions []models.Question, err error) {

	jsonFile, err := osOpen(path)

	if err != nil {
		return
	}

	defer jsonFile.Close()

	byteValue, err := readFile(jsonFile)
	if err != nil {
		return
	}

	json.Unmarshal(byteValue, &questions)

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
