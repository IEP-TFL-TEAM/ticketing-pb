package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `[
			{
				"id": "caza2oj0rtxw1gu",
				"created": "2024-08-02 04:24:16.239Z",
				"updated": "2024-11-04 08:30:36.169Z",
				"name": "tickets",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "svpre001",
						"name": "ticketNumber",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "g368qcuk",
						"name": "count",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "twyh6r6p",
						"name": "title",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "67jnb4rd",
						"name": "description",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "zndthmys",
						"name": "reportedBy",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "f0ojadtu",
						"name": "teamIds",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "rybge8g8z4jgrwi",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": null,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "1xmcaxpe",
						"name": "categoryId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "yrgvxehupjjixov",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "8gfahr20",
						"name": "categoryLevelId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "t0djc6rrp2vtymy",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "uct1dcix",
						"name": "teamEquipmentIds",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "4fgpahxwll5pzuj",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": null,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "qmffhsfk",
						"name": "regionId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "vvnj2k58yney9nn",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "hipbxxo3",
						"name": "areaId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "63z0ygwtpjs28a4",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "dygdueds",
						"name": "siteId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "ghtuohic73e7d35",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "adsoetyj",
						"name": "faultTypeId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "r8sn8obg2gksxgd",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "b5p1otrz",
						"name": "status",
						"type": "select",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"PENDING",
								"CLOSED"
							]
						}
					},
					{
						"system": false,
						"id": "eaxlfodc",
						"name": "closedBy",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "kaiuilip",
						"name": "attachment",
						"type": "file",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"mimeTypes": [],
							"thumbs": [],
							"maxSelect": 1,
							"maxSize": 5242880,
							"protected": true
						}
					},
					{
						"system": false,
						"id": "jmpzbaqp",
						"name": "closingRemarks",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "zktfi8or",
						"name": "technicianId",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "a3vplmoundlez9b",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "qxfuqxxw",
						"name": "cause",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "v0i45iaci8zjm58",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "hmepekdv",
						"name": "solution",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "u72l4ixeatr3dr6",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "edqlsddc",
						"name": "incidentStart",
						"type": "date",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "innrrgmg",
						"name": "incidentEnd",
						"type": "date",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "dv5ibxy3",
						"name": "slaStatus",
						"type": "select",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"SLA Met",
								"SLA Exceeded"
							]
						}
					},
					{
						"system": false,
						"id": "hm1wgftv",
						"name": "serviceImpact",
						"type": "select",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"Yes",
								"No"
							]
						}
					},
					{
						"system": false,
						"id": "7ru7nukv",
						"name": "servicesListIds",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "7n5u41jnd21jsnj",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": null,
							"displayFields": null
						}
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_FGKl1DR` + "`" + ` ON ` + "`" + `tickets` + "`" + ` (` + "`" + `ticketNumber` + "`" + `)"
				],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\"",
				"updateRule": "@request.auth.id != \"\"",
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "rybge8g8z4jgrwi",
				"created": "2024-08-05 12:11:19.785Z",
				"updated": "2024-10-30 04:01:46.120Z",
				"name": "teams",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "1vjku43s",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "vuvt4kkb",
						"name": "email",
						"type": "email",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"exceptDomains": [],
							"onlyDomains": []
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "5mec029rdzm42xd",
				"created": "2024-08-07 02:27:45.274Z",
				"updated": "2024-10-27 16:09:37.284Z",
				"name": "comments",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "j3q05lhd",
						"name": "content",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "czwi4zor",
						"name": "userId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "zoytsgvk",
						"name": "ticketId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "caza2oj0rtxw1gu",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "x57kgvpt",
						"name": "attachment",
						"type": "file",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"mimeTypes": [],
							"thumbs": [],
							"maxSelect": 1,
							"maxSize": 5242880,
							"protected": true
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\"",
				"updateRule": "@request.auth.id != \"\"",
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "jdq19kz788zjjg0",
				"created": "2024-08-07 02:31:10.804Z",
				"updated": "2024-11-03 14:19:04.005Z",
				"name": "history",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "96mpgiq9",
						"name": "action",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "xqxeca03",
						"name": "ticketId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "caza2oj0rtxw1gu",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\"",
				"updateRule": "@request.auth.id != \"\"",
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "l5qc8lbi3y6q2s7",
				"created": "2024-08-08 13:37:55.176Z",
				"updated": "2024-08-26 01:27:20.627Z",
				"name": "ticketCount",
				"type": "view",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "hnbwiyji",
						"name": "totalItems",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {
					"query": "SELECT\n  (ROW_NUMBER() OVER()) as id,\n  COUNT(*) as totalItems\nFROM tickets"
				}
			},
			{
				"id": "blal3gzv2ctzdg8",
				"created": "2024-08-08 13:51:27.906Z",
				"updated": "2024-08-27 02:50:57.157Z",
				"name": "location",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "qzxgvnpd",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "arllkzxo",
						"name": "latitude",
						"type": "number",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "kpt5atb0",
						"name": "longitude",
						"type": "number",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "yrgvxehupjjixov",
				"created": "2024-08-08 14:20:49.951Z",
				"updated": "2024-10-01 03:09:23.308Z",
				"name": "categories",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "y89bmbtb",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "ssdqmjdubsma48q",
				"created": "2024-08-08 16:21:50.495Z",
				"updated": "2024-08-26 01:34:10.590Z",
				"name": "equipment",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "3smuwr5g",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "e63ygltp",
						"name": "ticketId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "caza2oj0rtxw1gu",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "jjqfth5s",
						"name": "siteId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "ghtuohic73e7d35",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "t0djc6rrp2vtymy",
				"created": "2024-08-11 04:00:52.171Z",
				"updated": "2024-10-01 03:14:59.299Z",
				"name": "categorylevels",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "imcywn2h",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "3olljqmy",
						"name": "categoryId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "yrgvxehupjjixov",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "vvnj2k58yney9nn",
				"created": "2024-08-11 04:07:45.376Z",
				"updated": "2024-08-26 01:31:16.140Z",
				"name": "region",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "omwetar6",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "63z0ygwtpjs28a4",
				"created": "2024-08-11 04:15:27.072Z",
				"updated": "2024-08-26 01:35:52.462Z",
				"name": "area",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "lmznsftl",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "gldhdvzw",
						"name": "regionId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "vvnj2k58yney9nn",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "r8sn8obg2gksxgd",
				"created": "2024-08-11 14:49:38.064Z",
				"updated": "2024-08-26 01:33:50.577Z",
				"name": "faulttype",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "jfrlakan",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "ghtuohic73e7d35",
				"created": "2024-08-11 14:50:45.663Z",
				"updated": "2024-08-27 03:20:34.247Z",
				"name": "sites",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "azei6rql",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "jrqysjzm",
						"name": "areaId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "63z0ygwtpjs28a4",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "e4c5cdie",
						"name": "locationId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "blal3gzv2ctzdg8",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "4fgpahxwll5pzuj",
				"created": "2024-08-11 14:54:54.476Z",
				"updated": "2024-11-04 08:31:51.577Z",
				"name": "teamequipments",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "bledm9ty",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "5le8avmx",
						"name": "teamId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "rybge8g8z4jgrwi",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "c607t2ux8dbejkb",
				"created": "2024-08-11 15:06:31.457Z",
				"updated": "2024-10-28 14:37:55.595Z",
				"name": "recipients",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "5bo7aex0",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "uimhqz6z",
						"name": "email",
						"type": "email",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"exceptDomains": [],
							"onlyDomains": []
						}
					},
					{
						"system": false,
						"id": "5j1lbni9",
						"name": "verified",
						"type": "bool",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"deleteRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"options": {}
			},
			{
				"id": "67689ye74w7y9pa",
				"created": "2024-08-15 11:42:49.733Z",
				"updated": "2024-11-02 19:25:45.664Z",
				"name": "changerequests",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "y11exvrk",
						"name": "ticketNumber",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "uhhhh4fb",
						"name": "requestee",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "85gg9za10dpmgdy",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "wicsecye",
						"name": "title",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "elffmjsp",
						"name": "objective",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "v0w7jtth",
						"name": "startDate",
						"type": "date",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "cotbbonr",
						"name": "endDate",
						"type": "date",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "j8cox1ub",
						"name": "serviceImpact",
						"type": "select",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"Yes",
								"No"
							]
						}
					},
					{
						"system": false,
						"id": "izvmebzh",
						"name": "areaId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "63z0ygwtpjs28a4",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "rmwplino",
						"name": "regionId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "vvnj2k58yney9nn",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "sjeypqhd",
						"name": "siteId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "ghtuohic73e7d35",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "auiwdxui",
						"name": "involvedSystem",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "c51dwpud",
						"name": "teamIds",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "rybge8g8z4jgrwi",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": null,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "vjfh01kz",
						"name": "changeTeamId",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "yv3cl86s3yci1co",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "obm3pfks",
						"name": "summary",
						"type": "editor",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"convertUrls": false
						}
					},
					{
						"system": false,
						"id": "3vatrrma",
						"name": "servicesListIds",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "7n5u41jnd21jsnj",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": null,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "xap6irdp",
						"name": "awarenessToBeMade",
						"type": "select",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 5,
							"values": [
								"Yes",
								"No",
								"Internal",
								"External",
								"Media"
							]
						}
					},
					{
						"system": false,
						"id": "gv5cz7tu",
						"name": "attachment",
						"type": "file",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"mimeTypes": [],
							"thumbs": [],
							"maxSelect": 1,
							"maxSize": 5242880,
							"protected": true
						}
					},
					{
						"system": false,
						"id": "5poybqyk",
						"name": "submissionWithinFiveDays",
						"type": "select",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"Yes",
								"No"
							]
						}
					},
					{
						"system": false,
						"id": "1e3fnioc",
						"name": "durationAdhered",
						"type": "select",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"Yes",
								"No"
							]
						}
					},
					{
						"system": false,
						"id": "2nfd5s2e",
						"name": "serviceImpactCorrect",
						"type": "select",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"Yes",
								"No"
							]
						}
					},
					{
						"system": false,
						"id": "rhtukh9o",
						"name": "correctCustomerList",
						"type": "select",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"Yes",
								"No"
							]
						}
					},
					{
						"system": false,
						"id": "paz8ujeu",
						"name": "taskCompletion",
						"type": "select",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"Yes",
								"No",
								"Partially Completed"
							]
						}
					},
					{
						"system": false,
						"id": "34cuwl3h",
						"name": "closingRemarks",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "rixlzh0j",
						"name": "closingAttachment",
						"type": "file",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"mimeTypes": [],
							"thumbs": [],
							"maxSelect": 1,
							"maxSize": 5242880,
							"protected": true
						}
					},
					{
						"system": false,
						"id": "wqk3nun8",
						"name": "status",
						"type": "select",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"PENDING",
								"CLOSED"
							]
						}
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_IyHI72R` + "`" + ` ON ` + "`" + `changerequests` + "`" + ` (` + "`" + `ticketNumber` + "`" + `)"
				],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\"",
				"updateRule": "@request.auth.id != \"\"",
				"deleteRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"options": {}
			},
			{
				"id": "a3vplmoundlez9b",
				"created": "2024-08-19 22:48:51.727Z",
				"updated": "2024-08-26 01:43:15.015Z",
				"name": "technicians",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "csveqmuf",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "tb6typsh",
						"name": "teamId",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "rybge8g8z4jgrwi",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "w0xy9jxjhwjgnz3",
				"created": "2024-08-19 23:42:02.299Z",
				"updated": "2024-10-30 02:33:33.056Z",
				"name": "routinemaintenance",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "umrksdug",
						"name": "ticketNumber",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "jchr9asn",
						"name": "requestee",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "85gg9za10dpmgdy",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "xx4nl8ld",
						"name": "title",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "sdm0ora9",
						"name": "objective",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "1tnfqxdw",
						"name": "scopeOfWork",
						"type": "editor",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"convertUrls": false
						}
					},
					{
						"system": false,
						"id": "bqkrznt6",
						"name": "regionId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "vvnj2k58yney9nn",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "vi9iv3mh",
						"name": "areaId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "63z0ygwtpjs28a4",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "0rl6ali6",
						"name": "siteId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "ghtuohic73e7d35",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "xkyxcsf0",
						"name": "startDate",
						"type": "date",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "7aa7g2by",
						"name": "endDate",
						"type": "date",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "oztlvyel",
						"name": "serviceImpact",
						"type": "select",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"Yes",
								"No"
							]
						}
					},
					{
						"system": false,
						"id": "jcmaqvg5",
						"name": "duration",
						"type": "number",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": true
						}
					},
					{
						"system": false,
						"id": "wwitwnyi",
						"name": "teamIds",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "rybge8g8z4jgrwi",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": null,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "mprpgron",
						"name": "maintenanceTeamId",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "16udx7u72g9xwti",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "8ondgemt",
						"name": "servicesListIds",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "7n5u41jnd21jsnj",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": null,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "ec2bfeml",
						"name": "attachment",
						"type": "file",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"mimeTypes": [],
							"thumbs": [],
							"maxSelect": 1,
							"maxSize": 5242880,
							"protected": true
						}
					},
					{
						"system": false,
						"id": "2qtcz2em",
						"name": "taskCompletion",
						"type": "select",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"Yes",
								"No",
								"Partially Completed"
							]
						}
					},
					{
						"system": false,
						"id": "q9oklhkq",
						"name": "alarmsCleared",
						"type": "select",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"Yes",
								"No",
								"Alarms Outstanding"
							]
						}
					},
					{
						"system": false,
						"id": "5l2w8ift",
						"name": "serviceImpactCorrect",
						"type": "select",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"Yes",
								"No"
							]
						}
					},
					{
						"system": false,
						"id": "qslc84gg",
						"name": "closingRemarks",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "ndfvea8t",
						"name": "closingAttachment",
						"type": "file",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"mimeTypes": [],
							"thumbs": [],
							"maxSelect": 1,
							"maxSize": 5242880,
							"protected": true
						}
					},
					{
						"system": false,
						"id": "wlm2aq0f",
						"name": "status",
						"type": "select",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"PENDING",
								"CLOSED"
							]
						}
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_f57N5Vo` + "`" + ` ON ` + "`" + `routinemaintenance` + "`" + ` (` + "`" + `ticketNumber` + "`" + `)"
				],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\"",
				"updateRule": "@request.auth.id != \"\"",
				"deleteRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"options": {}
			},
			{
				"id": "v0i45iaci8zjm58",
				"created": "2024-08-20 00:11:03.836Z",
				"updated": "2024-10-09 05:06:54.186Z",
				"name": "causes",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "nrsn7vvl",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "7jcpdfk2",
						"name": "description",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "vm7jau3j",
						"name": "faultTypeId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "r8sn8obg2gksxgd",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "u72l4ixeatr3dr6",
				"created": "2024-08-20 00:11:42.122Z",
				"updated": "2024-10-09 05:07:01.349Z",
				"name": "solutions",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "gx8uwkpr",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "6wuy02xo",
						"name": "description",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "t5lyxrnr",
						"name": "faultTypeId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "r8sn8obg2gksxgd",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "_pb_users_auth_",
				"created": "2024-08-21 23:41:16.588Z",
				"updated": "2024-09-12 00:45:05.026Z",
				"name": "users",
				"type": "auth",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "users_name",
						"name": "firstName",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "mg7sg3k3",
						"name": "lastName",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "users_avatar",
						"name": "avatar",
						"type": "file",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"mimeTypes": [
								"image/jpeg",
								"image/png",
								"image/svg+xml",
								"image/gif",
								"image/webp"
							],
							"thumbs": null,
							"maxSelect": 1,
							"maxSize": 5242880,
							"protected": false
						}
					},
					{
						"system": false,
						"id": "yioss57m",
						"name": "role",
						"type": "select",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"admin",
								"staff"
							]
						}
					},
					{
						"system": false,
						"id": "g6nluuwo",
						"name": "teamId",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "rybge8g8z4jgrwi",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					}
				],
				"indexes": [],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": "",
				"deleteRule": null,
				"options": {
					"allowEmailAuth": true,
					"allowOAuth2Auth": true,
					"allowUsernameAuth": true,
					"exceptEmailDomains": null,
					"manageRule": "id = @request.auth.id && @request.auth.role = \"admin\"",
					"minPasswordLength": 8,
					"onlyEmailDomains": null,
					"onlyVerified": false,
					"requireEmail": false
				}
			},
			{
				"id": "7bfnibeo8zyr2si",
				"created": "2024-08-27 22:17:44.792Z",
				"updated": "2024-08-27 23:55:42.533Z",
				"name": "officelocations",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "vt4ynyu7",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "pnkui078",
						"name": "latitude",
						"type": "number",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "rfelrfbl",
						"name": "longitude",
						"type": "number",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "85gg9za10dpmgdy",
				"created": "2024-08-30 01:26:41.176Z",
				"updated": "2024-09-02 07:39:55.600Z",
				"name": "staff",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "udugues8",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "lbcqobt0",
						"name": "email",
						"type": "email",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"exceptDomains": [],
							"onlyDomains": []
						}
					},
					{
						"system": false,
						"id": "qaospocq",
						"name": "jobDescription",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "yv3cl86s3yci1co",
				"created": "2024-08-30 01:41:27.918Z",
				"updated": "2024-08-30 01:54:35.793Z",
				"name": "changeteams",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "flwuh8pt",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "6v2l2d2974asxtv",
				"created": "2024-08-30 01:44:28.488Z",
				"updated": "2024-09-02 01:33:31.565Z",
				"name": "changeteam_members",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "qz2bq2qr",
						"name": "changeTeamId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "yv3cl86s3yci1co",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "z7kn6ggv",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "p0axpg9c",
						"name": "isLead",
						"type": "bool",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "swowlfnn",
						"name": "jobDescription",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "16udx7u72g9xwti",
				"created": "2024-09-03 06:37:39.983Z",
				"updated": "2024-09-03 06:37:58.944Z",
				"name": "maintenanceteams",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "nutu3f3s",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "fjn5mhbky2ojf41",
				"created": "2024-09-03 22:48:23.455Z",
				"updated": "2024-09-03 22:48:23.455Z",
				"name": "maintenanceteam_members",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "ojch2eme",
						"name": "maintenanceTeamId",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "16udx7u72g9xwti",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "yjdjp8n8",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "9ezxp5n5",
						"name": "isLead",
						"type": "bool",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "p4ckhnxf",
						"name": "jobDescription",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "7n5u41jnd21jsnj",
				"created": "2024-10-22 23:31:14.506Z",
				"updated": "2024-10-22 23:31:14.506Z",
				"name": "serviceslist",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "m3hr522j",
						"name": "name",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id != \"\"",
				"viewRule": "@request.auth.id != \"\"",
				"createRule": "@request.auth.id != \"\"",
				"updateRule": "@request.auth.id != \"\"",
				"deleteRule": null,
				"options": {}
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		return nil
	})
}
