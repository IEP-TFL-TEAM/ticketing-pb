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
		jsonData := `{
			"id": "yv3cl86s3yci1co",
			"created": "2024-08-30 01:41:27.918Z",
			"updated": "2024-08-30 01:41:27.918Z",
			"name": "changeteam",
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
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("yv3cl86s3yci1co")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
