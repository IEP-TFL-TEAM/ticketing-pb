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
			"id": "c607t2ux8dbejkb",
			"created": "2024-08-11 15:06:31.457Z",
			"updated": "2024-08-11 15:06:31.457Z",
			"name": "stakeholders",
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

		collection, err := dao.FindCollectionByNameOrId("c607t2ux8dbejkb")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
