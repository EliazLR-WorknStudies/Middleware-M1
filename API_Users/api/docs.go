// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Justine Bachelard.",
            "email": "justine.bachelard@ext.uca.fr"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/users": {
            "get": {
                "description": "Get users.",
                "tags": [
                    "users"
                ],
                "summary": "Get users.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    },
                    "500": {
                        "description": "Something went wrong"
                    }
                }
            },
            "post": {
                "description": "Creates a user with username.",
                "tags": [
                    "users"
                ],
                "summary": "Creates a user with username.",
                "parameters": [
                    {
                        "description": "Username",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "422": {
                        "description": "Cannot parse username"
                    },
                    "500": {
                        "description": "Something went wrong"
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "Get a user.",
                "tags": [
                    "users"
                ],
                "summary": "Get a user.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User UUID formatted ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "422": {
                        "description": "Cannot parse id"
                    },
                    "500": {
                        "description": "Something went wrong"
                    }
                }
            },
            "put": {
                "description": "Update a user.",
                "tags": [
                    "users"
                ],
                "summary": "Update a user.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User UUID formatted ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User username",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Cannot parse id"
                    },
                    "500": {
                        "description": "Something went wrong"
                    }
                }
            },
            "delete": {
                "description": "Delete a user.",
                "tags": [
                    "users"
                ],
                "summary": "Delete a user.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User UUID formatted ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Cannot parse id"
                    },
                    "500": {
                        "description": "Something went wrong"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.User": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "middleware/users",
	Description:      "API to manage users.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
