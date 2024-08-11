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
			"id": "63z0ygwtpjs28a4",
			"created": "2024-08-11 04:15:27.072Z",
			"updated": "2024-08-11 04:15:27.072Z",
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
			"listRule": "",
			"viewRule": "",
			"createRule": "",
			"updateRule": "",
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

		collection, err := dao.FindCollectionByNameOrId("63z0ygwtpjs28a4")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
