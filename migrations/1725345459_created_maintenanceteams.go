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
			"id": "16udx7u72g9xwti",
			"created": "2024-09-03 06:37:39.983Z",
			"updated": "2024-09-03 06:37:39.983Z",
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
			"createRule": "@request.auth.id != \"\"",
			"updateRule": "@request.auth.id != \"\"",
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

		collection, err := dao.FindCollectionByNameOrId("16udx7u72g9xwti")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
