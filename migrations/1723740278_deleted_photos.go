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
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("5hkmlmpqsznb4t4")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	}, func(db dbx.Builder) error {
		jsonData := `{
			"id": "5hkmlmpqsznb4t4",
			"created": "2024-08-05 13:02:05.308Z",
			"updated": "2024-08-12 13:22:19.586Z",
			"name": "photos",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "9kqgwwui",
					"name": "attachment",
					"type": "file",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"mimeTypes": [
							"image/png",
							"image/vnd.mozilla.apng",
							"image/jpeg",
							"application/pdf",
							"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
							"application/msword"
						],
						"thumbs": [],
						"maxSelect": 1,
						"maxSize": 5242880,
						"protected": true
					}
				},
				{
					"system": false,
					"id": "3kkc3tll",
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
	})
}
