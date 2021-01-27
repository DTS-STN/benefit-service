package questions

import (
	"os"
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

// func osOpenMock(path string) (*os.File, error) {
// 	return os.Open(path)
// }

func loadQuestionsMock(lang string) (questions []models.Question, err error) {
	if lang == "fr" {
		return []models.Question{
			{
				ID:   1,
				Text: "Question 1 [FR]",
				Answers: []models.QuestionAnswers{
					{
						ID:   "1_fr",
						Text: "one [FR]",
					},
				},
			},
		}, nil
	}

	return []models.Question{
		{
			ID:   1,
			Text: "Question 1",
			Answers: []models.QuestionAnswers{
				{
					ID:   "1",
					Text: "one",
				},
			},
		},
	}, nil
}

// anything that should be run a the end of the unit tests should go here
func setupTests() {
	osOpen = os.Open
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
		assert.Equal(t, 1, len(data))
		assert.Equal(t, "Question 1", data[0].Text)
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
		assert.Equal(t, 1, len(data))
		assert.Equal(t, "Question 1 [FR]", data[0].Text)
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
		assert.Equal(t, 1, len(data))
		assert.Equal(t, "Question 1", data[0].Text)
	}
}

// func TestQuestionsNotEqual(t *testing.T) {
// 	setupTests()

// 	var realQuestionService = ServiceStruct{Filename: ""}

// 	// Expected result data
// 	expectedResult := []models.Question{{ID: "1", Description: "are you a resident of canada?", Answer: "", OpenFiscaIds: []string{"is_canadian_resident"}}}

// 	// Create a Mock for the interface
// 	qsMock := new(QuestionServiceMock)
// 	// Add a mock call request
// 	qsMock.On("LoadQuestions").
// 		Return([]models.Question{{ID: "2", Description: "are you a resident of canada?", Answer: "", OpenFiscaIds: []string{"2"}}}, nil)
// 	// Set the mock to be used by the code
// 	Service = QuestionInterface(qsMock)

// 	// redirect to test data
// 	osOpen = osOpenMock

// 	// Actual result data
// 	actual := realQuestionService.Questions()

// 	// Assertions
// 	assert.NotEqual(t, expectedResult, actual)
// }

// func TestPrefilledQuestions(t *testing.T) {
// 	setupTests()

// 	var realQuestionService = ServiceStruct{Filename: ""}

// 	// Expected result data
// 	expectedResult := []models.Question{{ID: "2", Description: "are you a resident of canada?", Answer: "", OpenFiscaIds: []string{"2"}}}

// 	// prefill test data
// 	questions = expectedResult

// 	// Create a Mock for the interface
// 	qsMock := new(QuestionServiceMock)
// 	// mock a different result from LoadQuestions, to prove that when questions is populated, LoadQuestions isn't called
// 	qsMock.On("LoadQuestions").
// 		Return([]models.Question{{ID: "1", Description: "are you a resident of canada?", Answer: "", OpenFiscaIds: []string{"1"}}}, nil)
// 	// Set the mock to be used by the code
// 	Service = QuestionInterface(qsMock)

// 	// redirect to test data
// 	osOpen = osOpenMock

// 	// Actual result data
// 	actual := realQuestionService.Questions()

// 	// Assertions
// 	assert.Equal(t, expectedResult, actual)
// 	assert.Equal(t, expectedResult, questions)
// }

// func TestReadFile(t *testing.T) {
// 	setupTests()

// 	var buffer bytes.Buffer
// 	buffer.WriteString("test read data. testing to see if readFile works")

// 	// expected results
// 	expected := buffer.Bytes()

// 	// actual results
// 	content, err := readFile(&buffer)

// 	// assertions
// 	assert.NoError(t, err)
// 	assert.Equal(t, expected, content)
// }

// func TestLoadQuestions(t *testing.T) {
// 	setupTests()

// 	// redirect to test data
// 	osOpen = osOpenMock

// 	// Expected result data
// 	expectedResult := []models.Question{{ID: "1", Description: "are you a resident?", Answer: "", OpenFiscaIds: []string{"is_canadian_resident"}}}

// 	// Actual result data
// 	actual, err := Service.LoadQuestions()

// 	// Assertions
// 	assert.NoError(t, err)
// 	assert.Equal(t, expectedResult, actual)
// }

// // Bug: overriding osOpen is not returning an error when the file is non existent
// func TestLoadQuestionsError(t *testing.T) {
// 	setupTests()

// 	// redirect to test data
// 	// want to return an error here
// 	osOpen = func(path string) (*os.File, error) {
// 		return &os.File{}, errors.New("missing file")
// 	}

// 	// Actual result data
// 	results, err := Service.LoadQuestions()

// 	// Assertions
// 	assert.Error(t, err)
// 	assert.Nil(t, results)
// }
