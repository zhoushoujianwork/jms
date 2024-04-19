// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/approval": {
            "post": {
                "description": "创建策略",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Approval"
                ],
                "summary": "创建审批策略",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/db.ApprovalMut"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/approval/:id": {
            "patch": {
                "description": "更新审批结果，可以是同意或者拒绝",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Approval"
                ],
                "summary": "更新审批",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "approval id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/db.ApprovalResult"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/key": {
            "get": {
                "description": "列出密钥，数据隐藏",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Key"
                ],
                "summary": "列出密钥",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.Key"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "添加密钥",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Key"
                ],
                "summary": "添加密钥",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "description": "key",
                        "name": "key",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/db.AddKeyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/key/:uuid": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Key"
                ],
                "summary": "删除密钥",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "key uuid",
                        "name": "uuid",
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
                    }
                }
            }
        },
        "/api/v1/login": {
            "post": {
                "description": "登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "user",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/policy": {
            "get": {
                "description": "获取策略列表，只能查某人或者某个组或者某个策略，不可组合查询，查用户会带出用户所在组的策略",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Policy"
                ],
                "summary": "获取策略列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "policy id",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "user",
                        "name": "user",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.Policy"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/policy/:id": {
            "put": {
                "description": "更新策略",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Policy"
                ],
                "summary": "更新策略",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "policy id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/db.PolicyMut"
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
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除策略",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Policy"
                ],
                "summary": "删除策略",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "policy id",
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
                    }
                }
            }
        },
        "/api/v1/profile": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "List profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "List profile",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.Profile"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "Create profile",
                "parameters": [
                    {
                        "description": "profile",
                        "name": "profile",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/db.CreateProfileRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/profile/:uuid": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "Update profile",
                "parameters": [
                    {
                        "description": "profile",
                        "name": "profile",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/db.CreateProfileRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "Delete profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "profile uuid",
                        "name": "uuid",
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
                    }
                }
            }
        },
        "/api/v1/user": {
            "get": {
                "description": "获取用户列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "获取用户列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "name 支持用户名或者email查询",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "group",
                        "name": "group",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.User"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/user/:id": {
            "put": {
                "description": "更新用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "更新用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/db.UserMut"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "支持数组会与现有组进行合并",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "追加用户组",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/db.UserPatchMut"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "db.Action": {
            "type": "string",
            "enum": [
                "connect",
                "deny_connect",
                "download",
                "deny_download",
                "upload",
                "deny_upload"
            ],
            "x-enum-varnames": [
                "Connect",
                "DenyConnect",
                "Download",
                "DenyDownload",
                "Upload",
                "DenyUpload"
            ]
        },
        "db.AddKeyRequest": {
            "type": "object",
            "required": [
                "key_id",
                "pem_base64"
            ],
            "properties": {
                "key_id": {
                    "description": "云上的key id，比如 skey-123456",
                    "type": "string"
                },
                "key_name": {
                    "description": "云上下载下来的名字，比如 jms-key.pem",
                    "type": "string"
                },
                "pem_base64": {
                    "description": "base64",
                    "type": "string"
                },
                "profile": {
                    "description": "云账号的 profile，比如 aws, aliyun",
                    "type": "string"
                }
            }
        },
        "db.ApprovalMut": {
            "type": "object",
            "required": [
                "applicant",
                "server_filter",
                "users"
            ],
            "properties": {
                "actions": {
                    "description": "申请动作，默认只有connect",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/db.Action"
                    }
                },
                "applicant": {
                    "description": "申请人AD名,或者email",
                    "type": "string"
                },
                "groups": {
                    "type": "array",
                    "items": {}
                },
                "name": {
                    "type": "string"
                },
                "period": {
                    "description": "审批周期，默认一周",
                    "allOf": [
                        {
                            "$ref": "#/definitions/db.Period"
                        }
                    ]
                },
                "server_filter": {
                    "$ref": "#/definitions/utils.ServerFilter"
                },
                "users": {
                    "type": "array",
                    "items": {}
                }
            }
        },
        "db.ApprovalResult": {
            "type": "object",
            "properties": {
                "applicant": {
                    "type": "string"
                },
                "is_pass": {
                    "type": "boolean"
                }
            }
        },
        "db.CreateProfileRequest": {
            "type": "object",
            "required": [
                "ak",
                "cloud",
                "name",
                "regions",
                "sk"
            ],
            "properties": {
                "ak": {
                    "type": "string"
                },
                "cloud": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "regions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "sk": {
                    "type": "string"
                }
            }
        },
        "db.Key": {
            "type": "object",
            "properties": {
                "isDelete": {
                    "type": "boolean"
                },
                "keyID": {
                    "type": "string"
                },
                "keyName": {
                    "type": "string"
                },
                "pemBase64": {
                    "type": "string"
                },
                "profile": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "db.Period": {
            "type": "string",
            "enum": [
                "1d",
                "1w",
                "1m",
                "1y",
                "ever"
            ],
            "x-enum-varnames": [
                "OneDay",
                "OneWeek",
                "OneMonth",
                "OneYear",
                "Forever"
            ]
        },
        "db.Policy": {
            "type": "object",
            "properties": {
                "actions": {
                    "type": "array",
                    "items": {}
                },
                "approval_id": {
                    "description": "审批ID",
                    "type": "string"
                },
                "approver": {
                    "description": "审批人",
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "expires_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_deleted": {
                    "type": "boolean"
                },
                "is_enabled": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "server_filter": {
                    "$ref": "#/definitions/utils.ServerFilter"
                },
                "updated_at": {
                    "type": "string"
                },
                "users": {
                    "type": "array",
                    "items": {}
                }
            }
        },
        "db.PolicyMut": {
            "type": "object",
            "properties": {
                "actions": {
                    "type": "array",
                    "items": {}
                },
                "expires_at": {
                    "type": "string"
                },
                "groups": {
                    "type": "array",
                    "items": {}
                },
                "is_enabled": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "server_filter": {
                    "$ref": "#/definitions/utils.ServerFilter"
                },
                "users": {
                    "type": "array",
                    "items": {}
                }
            }
        },
        "db.Profile": {
            "type": "object",
            "properties": {
                "ak": {
                    "type": "string"
                },
                "cloud": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "regions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "sk": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "db.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "dingtalk_dept_id": {
                    "type": "string"
                },
                "dingtalk_id": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "groups": {
                    "description": "组不在 jms维护这里只需要和机器 tag:Team 匹配即可。",
                    "type": "array",
                    "items": {}
                },
                "id": {
                    "type": "string"
                },
                "is_deleted": {
                    "type": "boolean"
                },
                "is_ldap": {
                    "type": "boolean"
                },
                "passwd": {
                    "description": "加密后的密码",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "db.UserMut": {
            "type": "object",
            "required": [
                "username"
            ],
            "properties": {
                "dingtalk_dept_id": {
                    "type": "string"
                },
                "dingtalk_id": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "groups": {
                    "type": "array",
                    "items": {}
                },
                "passwd": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "db.UserPatchMut": {
            "type": "object",
            "properties": {
                "groups": {
                    "type": "array",
                    "items": {}
                }
            }
        },
        "utils.ServerFilter": {
            "type": "object",
            "properties": {
                "env_type": {
                    "type": "string"
                },
                "ip_addr": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "team": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v1",
	Host:             "localhost:8013",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "cbs manager API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
