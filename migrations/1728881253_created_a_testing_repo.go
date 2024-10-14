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
			"id": "wyv8amrzqtq2mko",
			"created": "2024-10-14 04:47:33.350Z",
			"updated": "2024-10-14 04:47:33.350Z",
			"name": "a_testing_repo",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "hky4vrgj",
					"name": "field",
					"type": "text",
					"required": false,
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

		collection, err := dao.FindCollectionByNameOrId("wyv8amrzqtq2mko")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
