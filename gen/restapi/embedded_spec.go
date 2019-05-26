// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API description in Markdown.",
    "title": "Gremium API",
    "version": "0.1.0"
  },
  "host": "api.example.com",
  "basePath": "/api",
  "paths": {
    "/v1/quizes": {
      "get": {
        "description": "Optional extended description in Markdown.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "quiz"
        ],
        "summary": "Returns a list of quizes.",
        "operationId": "GetQuizes",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/quizQuestion"
              }
            }
          },
          "default": {
            "description": "Generic error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "quizQuestion": {
      "type": "object",
      "required": [
        "id",
        "question",
        "answers",
        "correct_answer_idx"
      ],
      "properties": {
        "answers": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "correct_answer_idx": {
          "type": "integer",
          "format": "int32"
        },
        "id": {
          "type": "string"
        },
        "question": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API description in Markdown.",
    "title": "Gremium API",
    "version": "0.1.0"
  },
  "host": "api.example.com",
  "basePath": "/api",
  "paths": {
    "/v1/quizes": {
      "get": {
        "description": "Optional extended description in Markdown.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "quiz"
        ],
        "summary": "Returns a list of quizes.",
        "operationId": "GetQuizes",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/quizQuestion"
              }
            }
          },
          "default": {
            "description": "Generic error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "quizQuestion": {
      "type": "object",
      "required": [
        "id",
        "question",
        "answers",
        "correct_answer_idx"
      ],
      "properties": {
        "answers": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "correct_answer_idx": {
          "type": "integer",
          "format": "int32"
        },
        "id": {
          "type": "string"
        },
        "question": {
          "type": "string"
        }
      }
    }
  }
}`))
}
