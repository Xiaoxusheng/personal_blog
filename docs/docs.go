// Code generated by swaggo/swag. DO NOT EDIT
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
        "/user/addarticle": {
            "post": {
                "description": "title category content token 为必填",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "添加文章接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "文章标题",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文章分类 0为技术类 1为生活类",
                        "name": "category",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文章内容 ",
                        "name": "content",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"添加成功！\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/articlelist": {
            "get": {
                "description": "token 为必填",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "获取文章接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "页数,不填默认为1",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "每页大小，默认为20",
                        "name": "number",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 200,\t\"data\": {\"length\": 4, \"list\": [\t{\t\"id\": 1, \"identification\": \"1068dcc0-4fd3-461d-9f4b-0100879457eb\",\t\"content\": \"1\",\t\"status\": 0,\t\"title\": \"1\",\t\"create_time\": \"1686055918562468\",\t\"ip\": \"127.0.0.1\", \"update_time\": \"1\", \"category\": \"1\" },]},\"msg\": \"获取数据成功！\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/deletearticle": {
            "delete": {
                "description": "identification token为必填",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "删除文章接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "文章唯一标识",
                        "name": "identification",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{ \"code\": 1, \"err\": \"删除文章不存在或者已经删除！\" }",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "用户名 密码 为必填",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "登录接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
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
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, \"msg\": \"登陆成功\",\"token\": \"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpbmRlbnRseSI6IjZhMmE0NjJjLWExMDctNDhlYS04MmU1LTc0ZTMwODMyN2U2ZiIsInVzZXJuYW1lIjoiYWRtaW4iLCJpc3MiOiJ0ZXN0IiwiZXhwIjoxNjc4Nzg2NTM1fQ.P4dJ_f2UGhKbpiIqHxTxghRKwKIlCpF2XjryHCSnKKk\"}",
                        "schema": {
                            "type": ""
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "用户名 密码 邮箱 为必填",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "注册接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{",
                        "schema": {
                            "type": ""
                        }
                    }
                }
            }
        },
        "/user/updatearticle": {
            "get": {
                "description": "identification content category title 为必填",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "更新文章接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "文章唯一标识",
                        "name": "identification",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文章内容",
                        "name": "content",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文章分类",
                        "name": "category",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文章标题",
                        "name": "title",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"msg\":\"添加成功！\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
