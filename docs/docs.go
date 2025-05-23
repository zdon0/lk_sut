// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Maks Mikhaylov",
            "url": "https://t.me/don101"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/user": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Add user",
                "parameters": [
                    {
                        "description": "user info",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/lk_sut_pkg_dto.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/lk_sut_pkg_dto.SimpleOkResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/lk_sut_pkg_dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/lk_sut_pkg_dto.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "description": "user info",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/lk_sut_pkg_dto.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/lk_sut_pkg_dto.SimpleOkResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/lk_sut_pkg_dto.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/lk_sut_pkg_dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/lk_sut_pkg_dto.ErrorResponse"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update password",
                "parameters": [
                    {
                        "description": "user info",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/lk_sut_pkg_dto.UpdateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/lk_sut_pkg_dto.SimpleOkResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/lk_sut_pkg_dto.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/lk_sut_pkg_dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/lk_sut_pkg_dto.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "lk_sut_pkg_dto.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "result": {
                    "type": "object"
                }
            }
        },
        "lk_sut_pkg_dto.SimpleOkResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "object"
                },
                "result": {
                    "$ref": "#/definitions/lk_sut_pkg_dto.SimpleOkResult"
                }
            }
        },
        "lk_sut_pkg_dto.SimpleOkResult": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "lk_sut_pkg_dto.UpdateUser": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string",
                    "example": "example@mail.com"
                },
                "new_password": {
                    "type": "string",
                    "example": "Password321"
                },
                "old_password": {
                    "type": "string",
                    "example": "Password123"
                }
            }
        },
        "lk_sut_pkg_dto.User": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string",
                    "example": "example@mail.com"
                },
                "password": {
                    "type": "string",
                    "example": "Password123"
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
	Title:            "Lk SUT Autocommitter",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
