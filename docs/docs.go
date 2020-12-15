// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/benefits": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Request Info on Benefits",
                "operationId": "benefits",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/renderings.BenefitsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/renderings.BenefitServiceError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/renderings.BenefitServiceError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/renderings.BenefitServiceError"
                        }
                    }
                }
            }
        },
        "/getquestions": {
            "get": {
                "description": "Returns question dependencies for requested benefits",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Returns question dependencies for requested benefits",
                "operationId": "GetBenefitQuestions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/renderings.BenefitQuestionsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/renderings.BenefitServiceError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/renderings.BenefitServiceError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/renderings.BenefitServiceError"
                        }
                    }
                }
            }
        },
        "/healthcheck": {
            "get": {
                "description": "Returns Healthy",
                "summary": "Returns Healthy",
                "operationId": "healthcheck",
                "responses": {
                    "200": {
                        "description": "Healthy",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/lifejourney": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Request Info on Life Journey",
                "operationId": "life-journey",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/renderings.LifeJourneyResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/renderings.BenefitServiceError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/renderings.BenefitServiceError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/renderings.BenefitServiceError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Benefits": {
            "type": "object",
            "properties": {
                "benefit_details": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.FieldDetails"
                    }
                },
                "description": {
                    "type": "string"
                },
                "description_fr": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "related_benefits": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
                    "type": "string"
                },
                "title_fr": {
                    "type": "string"
                }
            }
        },
        "models.FieldDetails": {
            "type": "object",
            "properties": {
                "field_long_description": {
                    "type": "string"
                },
                "field_long_description_fr": {
                    "type": "string"
                },
                "field_short_description": {
                    "type": "string"
                },
                "field_short_description_fr": {
                    "type": "string"
                },
                "fieldname": {
                    "type": "string"
                },
                "fieldname_fr": {
                    "type": "string"
                }
            }
        },
        "models.LifeJourney": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "description_fr": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lifejourney_details": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.FieldDetails"
                    }
                },
                "related_benefits": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
                    "type": "string"
                },
                "title_fr": {
                    "type": "string"
                }
            }
        },
        "models.Question": {
            "type": "object",
            "properties": {
                "answer": {
                    "type": "string"
                },
                "dependency_of": {
                    "description": "List of which benefits the question is a dependency of",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "openfisca_ids": {
                    "description": "The OpenFisca input variable name(s) attributed to this question",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "renderings.BenefitQuestionsResponse": {
            "type": "object",
            "properties": {
                "benefit_list": {
                    "description": "Array of specific benefits to get the questions for",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "life_journeys": {
                    "description": "Array of life journeys, which represent a subset of benefits",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "question_list": {
                    "description": "List of answered priority questions",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Question"
                    }
                },
                "request_date": {
                    "description": "Date period for request in ms since epoch",
                    "type": "integer"
                }
            }
        },
        "renderings.BenefitServiceError": {
            "type": "object",
            "properties": {
                "error_code": {
                    "type": "string"
                }
            }
        },
        "renderings.BenefitsResponse": {
            "type": "object",
            "properties": {
                "benefits": {
                    "description": "Life Journey ID",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Benefits"
                    }
                }
            }
        },
        "renderings.LifeJourneyResponse": {
            "type": "object",
            "properties": {
                "lifejourneys": {
                    "description": "Life Journey ID",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.LifeJourney"
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "https://benefit-service-dev.dev.dts-stn.com",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Benefit Service",
	Description: "This service returns information about Benefits",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
