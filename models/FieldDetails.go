package models

type FieldDetails struct {
	FieldName             string `json:"fieldname"`
	FieldShortDescription string `json:"field_short_description"`
	FieldLongDescription  string `json:"field_long_description"`
}
