// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/account": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "get account",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/account.AccountOutput"
                        }
                    }
                }
            }
        },
        "/flow": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "flow"
                ],
                "summary": "get flow list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/flow.FlowListObjectOutput"
                            }
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "flow"
                ],
                "summary": "create flow",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/flow.CreateFlowInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/flow.FlowDetailOutput"
                        }
                    }
                }
            }
        },
        "/flow/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "flow"
                ],
                "summary": "get flow detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "flow id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/flow.FlowDetailOutput"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "flow"
                ],
                "summary": "update flow",
                "parameters": [
                    {
                        "type": "string",
                        "description": "flow id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/flow.UpdateFlowInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/flow.FlowDetailOutput"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "flow"
                ],
                "summary": "delete flow",
                "parameters": [
                    {
                        "type": "string",
                        "description": "flow id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        }
    },
    "definitions": {
        "account.AccountOutput": {
            "type": "object",
            "properties": {
                "activated": {
                    "type": "boolean"
                },
                "membership": {
                    "type": "integer"
                }
            }
        },
        "flow.CreateFlowInput": {
            "type": "object",
            "properties": {
                "edges": {
                    "type": "string"
                },
                "nodes": {
                    "type": "string"
                },
                "owner": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "flow.FlowDetailOutput": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "edges": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nodes": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "flow.FlowListObjectOutput": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "flow.UpdateFlowInput": {
            "type": "object",
            "properties": {
                "edges": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nodes": {
                    "type": "string"
                },
                "owner": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Flow Editor",
	Description:      "Flow Editor APIs",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}