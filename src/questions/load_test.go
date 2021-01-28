package questions

import (
	"testing"

	"github.com/DTS-STN/benefit-service/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type QuestionServiceMock struct {
	mock.Mock
}

func (m *QuestionServiceMock) GetAll(lang string) ([]models.Question, error) {
	args := m.Called()
	return args.Get(0).([]models.Question), args.Error(1)
}

func (m *QuestionServiceMock) GetByID(lang, id string) (models.Question, error) {
	args := m.Called()
	return args.Get(0).(models.Question), args.Error(1)
}

func loadQuestionsMock(path string) (questions []models.Question, err error) {
	if path == "questions_fr.json" {
		return data_fr, nil
	}
	return data_en, nil
}

var data_en = []models.Question{
	{
		ID:   1,
		Text: "question 1",
		Answers: []models.QuestionAnswers{
			{
				ID:   "1_a",
				Text: "one_a",
			},
			{
				ID:   "1_b",
				Text: "one_b",
			},
		},
	},
	{
		ID:   2,
		Text: "question 2",
		Answers: []models.QuestionAnswers{
			{
				ID:   "2_a",
				Text: "two_a",
			},
			{
				ID:   "2_b",
				Text: "two_b",
			},
		},
	},
}

var data_fr = []models.Question{
	{
		ID:   1,
		Text: "question 1 [FR]",
		Answers: []models.QuestionAnswers{
			{
				ID:   "1_a [FR]",
				Text: "one_a",
			},
			{
				ID:   "1_b [FR]",
				Text: "one_b",
			},
		},
	},
	{
		ID:   2,
		Text: "question 2 [FR]",
		Answers: []models.QuestionAnswers{
			{
				ID:   "2_a [FR]",
				Text: "two_a",
			},
			{
				ID:   "2_b [FR]",
				Text: "two_b",
			},
		},
	},
}

// anything that should be run a the end of the unit tests should go here
func setupTests() {
	questionList = map[string][]models.Question{
		"en": {},
		"fr": {},
	}
	Service = new(ServiceStruct)
}

func TestGetAll_English(t *testing.T) {
	// setup tests
	setupTests()

	// mock unexported functions
	loadQuestions = loadQuestionsMock

	// result data
	data, err := Service.GetAll("en")

	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, 2, len(data))
		assert.Equal(t, "question 1", data[0].Text)
	}
}

func TestGetAll_French(t *testing.T) {
	// setup tests
	setupTests()

	// mock unexported functions
	loadQuestions = loadQuestionsMock

	// result data
	data, err := Service.GetAll("fr")

	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, 2, len(data))
		assert.Equal(t, "question 1 [FR]", data[0].Text)
	}
}

func TestGetAll_Default(t *testing.T) {
	// setup tests
	setupTests()

	// mock unexported functions
	loadQuestions = loadQuestionsMock

	// result data
	data, err := Service.GetAll("")

	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, 2, len(data))
		assert.Equal(t, "question 1", data[0].Text)
	}
}

func TestGetByID_English(t *testing.T) {
	var realService = ServiceStruct{}
	lang := "en"
	id := "1"

	// create a mock interface
	qsMock := new(QuestionServiceMock)

	// mock GetAll to return what we want
	qsMock.On("GetAll").Return(data_en, nil)
	// not sure this is right, but want to test the real function
	qsMock.On("GetByID").Return(realService.GetByID(lang, id))

	// set the service to the mock
	Service = QuestionInterface(qsMock)

	// this should call the real GetByID function, which will call the mock GetAll function
	result, err := Service.GetByID(lang, id)

	if assert.NoError(t, err) {
		assert.Equal(t, 1, result.ID)
		assert.Equal(t, "question 1", result.Text)
	}
}

func TestGetByID_French(t *testing.T) {
	var realService = ServiceStruct{}
	lang := "fr"
	id := "2"

	// create a mock interface
	qsMock := new(QuestionServiceMock)

	// mock GetAll to return what we want
	qsMock.On("GetAll").Return(data_fr, nil)
	// not sure this is right, but want to test the real function
	qsMock.On("GetByID").Return(realService.GetByID(lang, id))

	// set the service to the mock
	Service = QuestionInterface(qsMock)

	// this should call the real GetByID function, which will call the mock GetAll function
	result, err := Service.GetByID(lang, id)

	if assert.NoError(t, err) {
		assert.Equal(t, 2, result.ID)
		assert.Equal(t, "question 2 [FR]", result.Text)
	}
}
