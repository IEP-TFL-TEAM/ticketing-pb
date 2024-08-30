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
			"id": "6v2l2d2974asxtv",
			"created": "2024-08-30 01:44:28.488Z",
			"updated": "2024-08-30 01:44:28.488Z",
			"name": "changeteam_members",
			"type": "base",
			"system": false,
			"schema": [
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
					"required": true,
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
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("6v2l2d2974asxtv")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
