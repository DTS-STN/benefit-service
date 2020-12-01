package models

type FieldDetails struct {
	FieldName               string `json:"fieldname"`
	FieldNameFR             string `json:"fieldname_fr"`
	FieldShortDescription   string `json:"field_short_description"`
	FieldShortDescriptionFr string `json:"field_short_description_fr"`
	FieldLongDescription    string `json:"field_long_description"`
	FieldLongDescriptionFr  string `json:"field_long_description_fr"`
}
