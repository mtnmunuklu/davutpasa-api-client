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
        "/IAlertAct/GetActionDefinitions": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Search for action definitions based on a query.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ActionDefinations"
                ],
                "summary": "Search Action Definitions",
                "parameters": [
                    {
                        "description": "Search Action Definitions Request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api_davudpasha.ActionDefinationsSearchRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api_davudpasha.ActionDefinationsSearchResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api_davudpasha.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/api_davudpasha.ErrorResponse"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests",
                        "schema": {
                            "$ref": "#/definitions/api_davudpasha.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/IAssetAct/GetAllAsset2": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieve assets that match a search filter.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Assets"
                ],
                "summary": "Search Assets",
                "parameters": [
                    {
                        "description": "Search Assets Request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api_davudpasha.AssetsSearchRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/api_davudpasha.AssetsSearchResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api_davudpasha.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/api_davudpasha.ErrorResponse"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests",
                        "schema": {
                            "$ref": "#/definitions/api_davudpasha.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api_davudpasha.Action": {
            "type": "object",
            "properties": {
                "ActionParameters": {
                    "description": "Parameters of the action.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api_davudpasha.ActionParameter"
                    }
                },
                "ActionRefId": {
                    "description": "Reference ID for the action.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/github_com_mtnmunuklu_davudpasha-api-client-go_api_common.NullableString"
                        }
                    ]
                },
                "ActionType": {
                    "description": "Type of the action.",
                    "type": "string"
                },
                "Data": {
                    "description": "Data associated with the action.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/github_com_mtnmunuklu_davudpasha-api-client-go_api_common.NullableString"
                        }
                    ]
                },
                "additionalProperties": {
                    "description": "Additional properties not defined in the struct.",
                    "type": "object",
                    "additionalProperties": true
                }
            }
        },
        "api_davudpasha.ActionDefinationsSearchRequest": {
            "type": "object",
            "properties": {
                "additionalProperties": {
                    "description": "Additional properties not defined in the struct.",
                    "type": "object",
                    "additionalProperties": true
                },
                "application": {
                    "description": "Specifies the application for which the action definitions are being searched.",
                    "type": "string"
                },
                "searchFilter": {
                    "description": "Filter for searching action definations.",
                    "type": "string"
                },
                "smartRestRequestContext": {
                    "description": "Context for the Smart REST request.",
                    "type": "string"
                }
            }
        },
        "api_davudpasha.ActionDefinationsSearchResponse": {
            "type": "object",
            "properties": {
                "Action": {
                    "description": "Contains the details of the action.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/api_davudpasha.Action"
                        }
                    ]
                },
                "ActionDefinitionId": {
                    "description": "Unique identifier for the action definition.",
                    "type": "string"
                },
                "Application": {
                    "description": "Specifies the application to which the action definition belongs.",
                    "type": "string"
                },
                "Author": {
                    "description": "The author of the action definition.",
                    "type": "string"
                },
                "Description": {
                    "description": "Description of the action definition, which can be null.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/github_com_mtnmunuklu_davudpasha-api-client-go_api_common.NullableString"
                        }
                    ]
                },
                "IsPublic": {
                    "description": "Indicates whether the action definition is public.",
                    "type": "boolean"
                },
                "Name": {
                    "description": "Name of the action definition.",
                    "type": "string"
                },
                "Type": {
                    "description": "Type of the action definition.",
                    "type": "string"
                },
                "additionalProperties": {
                    "description": "Additional properties not defined in the struct.",
                    "type": "object",
                    "additionalProperties": true
                }
            }
        },
        "api_davudpasha.ActionParameter": {
            "type": "object",
            "properties": {
                "Key": {
                    "description": "Key of the action parameter.",
                    "type": "string"
                },
                "Value": {
                    "description": "Value of the action parameter.",
                    "type": "string"
                },
                "additionalProperties": {
                    "description": "Additional properties not defined in the struct.",
                    "type": "object",
                    "additionalProperties": true
                }
            }
        },
        "api_davudpasha.AssetsAsset": {
            "type": "object",
            "properties": {
                "AssetId": {
                    "description": "Unique identifier for the asset.",
                    "type": "string"
                },
                "AssetType": {
                    "description": "Type of the asset.",
                    "type": "string"
                },
                "Credential": {
                    "description": "Credential information for the asset, which can be null.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/github_com_mtnmunuklu_davudpasha-api-client-go_api_common.NullableString"
                        }
                    ]
                },
                "CustomerId": {
                    "description": "Identifier for the customer associated with the asset.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/github_com_mtnmunuklu_davudpasha-api-client-go_api_common.NullableString"
                        }
                    ]
                },
                "DistinguishedName": {
                    "description": "Distinguished name of the asset.",
                    "type": "string"
                },
                "EnterDate": {
                    "description": "Date when the asset was entered into the system.",
                    "type": "string"
                },
                "GroupId": {
                    "description": "Identifier for the group to which the asset belongs.",
                    "type": "string"
                },
                "HostName": {
                    "description": "Hostname of the asset.",
                    "type": "string"
                },
                "InventoryInfo": {
                    "description": "Inventory information for the asset.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/api_davudpasha.AssetsInventoryInfo"
                        }
                    ]
                },
                "Ip": {
                    "description": "IP address of the asset.",
                    "type": "string"
                },
                "IsDeleted": {
                    "description": "Indicates whether the asset is deleted.",
                    "type": "boolean"
                },
                "IsManage": {
                    "description": "Indicates whether the asset is manageable.",
                    "type": "boolean"
                },
                "IsManuel": {
                    "description": "Indicates whether the asset was manually added.",
                    "type": "boolean"
                },
                "IsPassive": {
                    "description": "Indicates whether the asset is passive.",
                    "type": "boolean"
                },
                "LastAccessDate": {
                    "description": "Date of the last access to the asset.",
                    "type": "string"
                },
                "LastUserOfSession": {
                    "description": "Last user who accessed the asset's session.",
                    "type": "string"
                },
                "MacAddress": {
                    "description": "MAC address of the asset.",
                    "type": "string"
                },
                "Name": {
                    "description": "Name of the asset.",
                    "type": "string"
                },
                "ObjectGUID": {
                    "description": "Unique object GUID for the asset.",
                    "type": "string"
                },
                "OsInfo": {
                    "description": "Information about the operating system of the asset.",
                    "type": "string"
                },
                "ParentGroupIds": {
                    "description": "List of parent group IDs to which the asset belongs.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/github_com_mtnmunuklu_davudpasha-api-client-go_api_common.NullableList-string"
                        }
                    ]
                },
                "Tag": {
                    "description": "List of tags associated with the asset.",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "TagData": {
                    "description": "Additional tag data, which can be null.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/github_com_mtnmunuklu_davudpasha-api-client-go_api_common.NullableString"
                        }
                    ]
                },
                "TenantID": {
                    "description": "Unique identifier for the tenant.",
                    "type": "string"
                },
                "Timestamp": {
                    "description": "Timestamp of when the asset was created or last updated.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/github_com_mtnmunuklu_davudpasha-api-client-go_api_common.NullableString"
                        }
                    ]
                },
                "_id": {
                    "description": "Unique identifier for the asset.",
                    "type": "string"
                },
                "additionalProperties": {
                    "description": "Additional properties not defined in the struct.",
                    "type": "object",
                    "additionalProperties": true
                }
            }
        },
        "api_davudpasha.AssetsInventoryInfo": {
            "type": "object",
            "properties": {
                "OperationSystemInfos": {
                    "description": "Information about the operating system of the asset.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/api_davudpasha.AssetsOperationSystemInfo"
                        }
                    ]
                },
                "additionalProperties": {
                    "description": "Additional properties not defined in the struct.",
                    "type": "object",
                    "additionalProperties": true
                }
            }
        },
        "api_davudpasha.AssetsOperationSystemInfo": {
            "type": "object",
            "properties": {
                "BuildNumber": {
                    "description": "Build number of the operating system.",
                    "type": "string"
                },
                "Caption": {
                    "description": "Caption or name of the operating system.",
                    "type": "string"
                },
                "Description": {
                    "description": "Description of the operating system.",
                    "type": "string"
                },
                "Manufacturer": {
                    "description": "Manufacturer of the operating system.",
                    "type": "string"
                },
                "OSArchitecture": {
                    "description": "Architecture of the operating system (e.g., 32-bit, 64-bit).",
                    "type": "string"
                },
                "SerialNumber": {
                    "description": "Serial number of the operating system.",
                    "type": "string"
                },
                "ServisPack": {
                    "description": "Service pack version of the operating system.",
                    "type": "string"
                },
                "Version": {
                    "description": "Version of the operating system.",
                    "type": "string"
                },
                "additionalProperties": {
                    "description": "Additional properties not defined in the struct.",
                    "type": "object",
                    "additionalProperties": true
                }
            }
        },
        "api_davudpasha.AssetsSearchRequest": {
            "type": "object",
            "properties": {
                "additionalProperties": {
                    "description": "Additional properties not defined in the struct.",
                    "type": "object",
                    "additionalProperties": true
                },
                "pageNumber": {
                    "description": "Page number for pagination.",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "Number of items per page for pagination.",
                    "type": "integer"
                },
                "searchFilter": {
                    "description": "Primary filter for searching assets.",
                    "type": "string"
                },
                "searchFilter2": {
                    "description": "Secondary filter for searching assets.",
                    "type": "string"
                },
                "smartRestRequestContext": {
                    "description": "Context for the Smart REST request.",
                    "type": "string"
                }
            }
        },
        "api_davudpasha.AssetsSearchResponse": {
            "type": "object",
            "properties": {
                "additionalProperties": {
                    "description": "Additional properties not defined in the struct.",
                    "type": "object",
                    "additionalProperties": true
                },
                "assets": {
                    "description": "List of assets matching the search criteria.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api_davudpasha.AssetsAsset"
                    }
                },
                "count": {
                    "description": "Total count of assets matching the search criteria.",
                    "type": "integer"
                }
            }
        },
        "api_davudpasha.ErrorResponse": {
            "type": "object",
            "properties": {
                "additionalProperties": {
                    "description": "Additional properties not defined in the struct.",
                    "type": "object",
                    "additionalProperties": true
                },
                "errors": {
                    "description": "A list of errors",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "github_com_mtnmunuklu_davudpasha-api-client-go_api_common.NullableList-string": {
            "type": "object"
        },
        "github_com_mtnmunuklu_davudpasha-api-client-go_api_common.NullableString": {
            "type": "object"
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
