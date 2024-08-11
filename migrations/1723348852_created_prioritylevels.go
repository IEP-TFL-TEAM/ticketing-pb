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
			"id": "t0djc6rrp2vtymy",
			"created": "2024-08-11 04:00:52.171Z",
			"updated": "2024-08-11 04:00:52.171Z",
			"name": "prioritylevels",
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
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
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

		collection, err := dao.FindCollectionByNameOrId("t0djc6rrp2vtymy")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
