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
        "/task": {
            "get": {
                "description": "Retrieve all tasks",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Get all tasks",
                "responses": {
                    "200": {
                        "description": "List of tasks",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Task"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/shared.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new task, body required",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Create a new task",
                "parameters": [
                    {
                        "description": "Task object",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.TaskInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Task created successfully",
                        "schema": {
                            "$ref": "#/definitions/domain.Task"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/shared.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity",
                        "schema": {
                            "$ref": "#/definitions/shared.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/task/{id}": {
            "get": {
                "description": "Retrieve a specific task using its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Get a task by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task details",
                        "schema": {
                            "$ref": "#/definitions/domain.Task"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/shared.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity",
                        "schema": {
                            "$ref": "#/definitions/shared.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing task by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Update a task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated task data",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.TaskInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task updated successfully",
                        "schema": {
                            "$ref": "#/definitions/domain.Task"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/shared.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity",
                        "schema": {
                            "$ref": "#/definitions/shared.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a task by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Delete a task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/shared.IDResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/shared.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable entity",
                        "schema": {
                            "$ref": "#/definitions/shared.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Task": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "domain.TaskInput": {
            "type": "object",
            "required": [
                "status",
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "status": {
                    "type": "string",
                    "enum": [
                        "new",
                        "in_progress",
                        "done"
                    ]
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "shared.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "shared.IDResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Tasks",
	Description:      "These are docs for test assignment from Skills Rock.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
