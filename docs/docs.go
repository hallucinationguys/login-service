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
        "/auth/login": {
            "post": {
                "description": "login user, returns user and set session",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Login new user",
                "parameters": [
                    {
                        "description": "Login user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usermodel.LoginUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usermodel.LoginUserResponse"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/common.AppError"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Register new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Register new user",
                "parameters": [
                    {
                        "description": "Login user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usermodel.UserCreate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/common.successRes"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/common.AppError"
                        }
                    }
                }
            }
        },
        "/profile": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get user info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "Profile user",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.successRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.AppError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/common.AppError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.AppError": {
            "type": "object",
            "properties": {
                "error_key": {
                    "type": "string"
                },
                "log": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "common.successRes": {
            "type": "object",
            "properties": {
                "data": {},
                "filter": {},
                "paging": {},
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "usermodel.LoginUserRequest": {
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
        "usermodel.LoginUserResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "expired_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/usermodel.UserResponse"
                }
            }
        },
        "usermodel.UserCreate": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/usermodel.UserRole"
                },
                "status": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "usermodel.UserResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/usermodel.UserRole"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "usermodel.UserRole": {
            "type": "integer",
            "enum": [
                1,
                2
            ],
            "x-enum-varnames": [
                "RoleUser",
                "RoleAdmin"
            ]
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Login Service API",
	Description:      "Ecosystem Hallucination Guys API Document",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
