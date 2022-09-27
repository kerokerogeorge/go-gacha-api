// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/character": {
            "post": {
                "description": "新しいキャラクターを作成します",
                "consumes": [
                    "application/json"
                ],
                "summary": "キャラクターを作成するAPI",
                "parameters": [
                    {
                        "description": "name",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "imgUrl",
                        "name": "imgUrl",
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
                            "$ref": "#/definitions/model.Character"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.Error"
                        }
                    }
                }
            }
        },
        "/character/emmition_rates": {
            "get": {
                "description": "キャラクター一覧を排出率とともに取得します",
                "consumes": [
                    "application/json"
                ],
                "summary": "キャラクター一覧を排出率とともに取得するAPI",
                "parameters": [
                    {
                        "type": "string",
                        "description": "gachaId",
                        "name": "gachaId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.CharacterWithEmmitionRate"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.Error"
                        }
                    }
                }
            }
        },
        "/character/list": {
            "get": {
                "description": "登録されているキャラクター一覧を取得します",
                "consumes": [
                    "application/json"
                ],
                "summary": "キャラクター一覧を取得するAPI",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Character"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.Error"
                        }
                    }
                }
            }
        },
        "/character/{characterId}": {
            "delete": {
                "description": "キャラクターを一件削除します",
                "consumes": [
                    "application/json"
                ],
                "summary": "キャラクターを削除するAPI",
                "parameters": [
                    {
                        "type": "string",
                        "description": "characterId",
                        "name": "characterId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.Error"
                        }
                    }
                }
            }
        },
        "/gacha": {
            "post": {
                "description": "新しいガチャを作成し、排出率をキャラクターに割り当てます",
                "consumes": [
                    "application/json"
                ],
                "summary": "新しいガチャを作成するAPI",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.CreateGachaResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.Error"
                        }
                    }
                }
            }
        },
        "/gacha/draw/{gachaId}": {
            "post": {
                "description": "ガチャを実行し、キャラクターを取得します",
                "consumes": [
                    "application/json"
                ],
                "summary": "ガチャを実行するAPI",
                "parameters": [
                    {
                        "type": "string",
                        "description": "x-token",
                        "name": "x-token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "gachaId",
                        "name": "gachaId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "ガチャを実行する回数",
                        "name": "times",
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
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Result"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.Error"
                        }
                    }
                }
            }
        },
        "/gacha/list": {
            "get": {
                "description": "ガチャ一覧を取得します",
                "consumes": [
                    "application/json"
                ],
                "summary": "ガチャ一覧を取得するAPI",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handler.GachaListResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.Error"
                        }
                    }
                }
            }
        },
        "/gacha/{gachaId}": {
            "get": {
                "description": "新しいガチャと登録されているキャラクターの排出率を取得する",
                "consumes": [
                    "application/json"
                ],
                "summary": "ガチャを一件取得するAPI",
                "parameters": [
                    {
                        "type": "string",
                        "description": "gachaId",
                        "name": "gachaId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Gacha"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "ガチャを一件削除します",
                "consumes": [
                    "application/json"
                ],
                "summary": "ガチャを削除するAPI",
                "parameters": [
                    {
                        "type": "string",
                        "description": "gachaId",
                        "name": "gachaId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.Error"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "ユーザーを一件取得する",
                "consumes": [
                    "application/json"
                ],
                "summary": "新しいユーザーを一件取得するAPI",
                "parameters": [
                    {
                        "type": "string",
                        "description": "x-token",
                        "name": "x-token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.GetUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.Error"
                        }
                    }
                }
            },
            "put": {
                "description": "ユーザーを一件更新する",
                "consumes": [
                    "application/json"
                ],
                "summary": "ユーザー情報を更新するAPI",
                "parameters": [
                    {
                        "type": "string",
                        "description": "x-token",
                        "name": "x-token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "ユーザー名",
                        "name": "name",
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
                            "$ref": "#/definitions/handler.UpdateUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "新しいユーザーを作成します",
                "consumes": [
                    "application/json"
                ],
                "summary": "新しいユーザーを作成するAPI",
                "parameters": [
                    {
                        "description": "name",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "address",
                        "name": "address",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/handler.CreateUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "ユーザーを一件削除する",
                "consumes": [
                    "application/json"
                ],
                "summary": "ユーザー情報を削除するAPI",
                "parameters": [
                    {
                        "type": "string",
                        "description": "x-token",
                        "name": "x-token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.Error"
                        }
                    }
                }
            }
        },
        "/user/characters": {
            "get": {
                "description": "ユーザー所持キャラクター一覧を取得します",
                "consumes": [
                    "application/json"
                ],
                "summary": "ユーザー所持キャラクター一覧を取得するAPI",
                "parameters": [
                    {
                        "type": "string",
                        "description": "x-token",
                        "name": "x-token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Result"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.Error"
                        }
                    }
                }
            }
        },
        "/user/list": {
            "get": {
                "description": "ユーザー一覧を取得します",
                "consumes": [
                    "application/json"
                ],
                "summary": "ユーザー一覧を取得するAPI",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.UserListResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.CreateGachaResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "handler.CreateUserResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "handler.GachaListResponse": {
            "type": "object",
            "properties": {
                "gachaId": {
                    "type": "string"
                }
            }
        },
        "handler.GetUserResponse": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "handler.UpdateUserResponse": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "handler.UserListResponse": {
            "type": "object",
            "properties": {
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.User"
                    }
                }
            }
        },
        "helper.Error": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.Character": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "imgUrl": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "model.CharacterWithEmmitionRate": {
            "type": "object",
            "properties": {
                "characterId": {
                    "type": "string"
                },
                "emissionRate": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.Gacha": {
            "type": "object",
            "properties": {
                "characters": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.CharacterWithEmmitionRate"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "model.Result": {
            "type": "object",
            "properties": {
                "characterId": {
                    "type": "string"
                },
                "emissionRate": {
                    "type": "number"
                },
                "imgUrl": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "userCharacterId": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:8000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Gacha-API Docs",
	Description:      "ガチャアプリのAPI仕様書です",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
