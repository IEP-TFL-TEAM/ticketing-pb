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
			"id": "l5qc8lbi3y6q2s7",
			"created": "2024-08-08 13:37:55.176Z",
			"updated": "2024-08-08 13:37:55.176Z",
			"name": "ticketCount",
			"type": "view",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "g1qihvwx",
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
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {
				"query": "SELECT\n  (ROW_NUMBER() OVER()) as id,\n  COUNT(*) as totalItems\nFROM tickets"
			}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("l5qc8lbi3y6q2s7")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
