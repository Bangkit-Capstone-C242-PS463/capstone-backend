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
        "/v1/auth/login": {
            "post": {
                "description": "Authenticates a user and returns a data upon successful login.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "User login credentials",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login success, returns user data",
                        "schema": {
                            "$ref": "#/definitions/dto.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request - Invalid input",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized - Invalid credentials",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found - User not found",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/auth/signup": {
            "post": {
                "description": "Registers a new user in the system.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Sign up a new user",
                "parameters": [
                    {
                        "description": "User registration details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Sign-up success status",
                        "schema": {
                            "$ref": "#/definitions/dto.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request - Invalid input",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict - User already exists",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Sign-up failed",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/health": {
            "get": {
                "description": "Returns the current health status of the server.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Get server health status",
                "responses": {
                    "200": {
                        "description": "Server is healthy",
                        "schema": {
                            "$ref": "#/definitions/dto.ServerHealthResponse"
                        }
                    }
                }
            }
        },
        "/v1/user/history": {
            "get": {
                "description": "Retrieves the history of user diagnoses based on their submitted symptoms.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "history"
                ],
                "summary": "Get user diagnosis history",
                "responses": {
                    "200": {
                        "description": "Diagnosis history retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/dto.GetAllHistoryResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Failed to fetch history",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a specific diagnosis history record by its ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "history"
                ],
                "summary": "Delete a diagnosis history record",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the history record to delete",
                        "name": "history_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Deletion successful",
                        "schema": {
                            "$ref": "#/definitions/dto.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request - Missing history_id",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Failed to delete history",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "dto.GetAllHistoryResponse": {
            "type": "object",
            "properties": {
                "history": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.History"
                    }
                }
            }
        },
        "dto.History": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "result": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "dto.LoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.LoginResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dto.ServerHealthResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "example": "healthy"
                }
            }
        },
        "dto.SignUpRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.SuccessResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "string",
                    "example": "true"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}