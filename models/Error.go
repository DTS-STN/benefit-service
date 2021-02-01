package models

type Error struct {
	Status       int    `json:"status_code"`
	ErrorMessage string `json:"error_message"`
}
