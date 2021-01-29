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
                "parameters": [
                    {
                        "type": "string",
                        "description": "The language the response should be in. Defaults to English. English and French supported.",
                        "name": "lang",
                        "in": "query"
                    }
                ],
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
        "/benefits/apply": {
            "post": {
                "security": [
                    {
                        "OAuth2AccessCode": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Request redirect url to IEG for a particular benefit",
                "operationId": "benefits-apply",
                "parameters": [
                    {
                        "description": "the benefit you are requesting an apply redirect for",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/bindings.BenefitApplyRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "the bearer token for a particular user",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "307": {
                        "description": "Temporary Redirect",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
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
        "/questions": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get a list of questions for pre-screening eligibilty",
                "operationId": "questions",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The language the response should be in. Defaults to English. English and French supported.",
                        "name": "lang",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/renderings.QuestionResponse"
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
        "bindings.BenefitApplyRequest": {
            "type": "object",
            "properties": {
                "benefitType": {
                    "type": "string"
                },
                "incomeDetails": {
                    "type": "string"
                },
                "outOfWork": {
                    "type": "string"
                },
                "reasonForSeperation": {
                    "type": "string"
                },
                "regularLookingForWork": {
                    "type": "string"
                }
            }
        },
        "models.Benefits": {
            "type": "object",
            "properties": {
                "api_url": {
                    "type": "string"
                },
                "benefit_key": {
                    "type": "string"
                },
                "benefit_tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "benefit_type": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "long_description": {
                    "type": "string"
                },
                "redirect_url": {
                    "type": "string"
                },
                "related_benefits": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "service_type": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.Error": {
            "type": "object",
            "properties": {
                "errorMessage": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "models.Question": {
            "type": "object",
            "properties": {
                "answers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.QuestionAnswers"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "models.QuestionAnswers": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
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
                "benefit": {
                    "$ref": "#/definitions/models.Benefits"
                },
                "benefits": {
                    "description": "Life Journey ID",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Benefits"
                    }
                }
            }
        },
        "renderings.QuestionResponse": {
            "type": "object",
            "properties": {
                "question": {
                    "$ref": "#/definitions/models.Question"
                },
                "questions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Question"
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
