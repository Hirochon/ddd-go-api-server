// Code generated by go-swagger; DO NOT EDIT.

package restapiv1

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
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "- API定義書\n",
    "title": "ddd-go-api-server",
    "version": "0.1.0"
  },
  "host": "localhost:8080",
  "basePath": "/v1/",
  "paths": {
    "/{id}/status": {
      "post": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Hoge"
        ],
        "summary": "ステータスを作成する",
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "object": {
                  "type": "object",
                  "properties": {
                    "id": {
                      "type": "integer",
                      "example": 2
                    }
                  }
                },
                "title": {
                  "type": "string",
                  "example": "迷走中"
                }
              }
            }
          }
        ],
        "responses": {
          "201": {
            "description": "成功",
            "schema": {
              "type": "object",
              "properties": {
                "id": {
                  "type": "integer",
                  "example": 2
                },
                "title": {
                  "type": "string",
                  "example": "迷走中"
                }
              }
            }
          },
          "400": {
            "description": "失敗",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string",
                  "example": "Bad Request"
                }
              }
            }
          }
        }
      }
    }
  },
  "tags": [
    {
      "description": "hogehogeするとこ",
      "name": "Hoge"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "- API定義書\n",
    "title": "ddd-go-api-server",
    "version": "0.1.0"
  },
  "host": "localhost:8080",
  "basePath": "/v1/",
  "paths": {
    "/{id}/status": {
      "post": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Hoge"
        ],
        "summary": "ステータスを作成する",
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "object": {
                  "type": "object",
                  "properties": {
                    "id": {
                      "type": "integer",
                      "example": 2
                    }
                  }
                },
                "title": {
                  "type": "string",
                  "example": "迷走中"
                }
              }
            }
          }
        ],
        "responses": {
          "201": {
            "description": "成功",
            "schema": {
              "type": "object",
              "properties": {
                "id": {
                  "type": "integer",
                  "example": 2
                },
                "title": {
                  "type": "string",
                  "example": "迷走中"
                }
              }
            }
          },
          "400": {
            "description": "失敗",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string",
                  "example": "Bad Request"
                }
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "PostIDStatusParamsBodyObject": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "example": 2
        }
      }
    }
  },
  "tags": [
    {
      "description": "hogehogeするとこ",
      "name": "Hoge"
    }
  ]
}`))
}
