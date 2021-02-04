package models

// Question struct used for benefit dependency information
type Question struct {
	ID      int               `json:"id"`
	Text    string            `json:"text"`
	Value   string            `json:"value"`
	Answers []QuestionAnswers `json:"answers"`
}

// QuestionAnswers is the the radio button options for a question
type QuestionAnswers struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}
