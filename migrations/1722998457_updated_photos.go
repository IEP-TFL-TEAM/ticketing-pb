package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("5hkmlmpqsznb4t4")
		if err != nil {
			return err
		}

		// update
		edit_attachment := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
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
					"image/webp"
				],
				"thumbs": [],
				"maxSelect": 1,
				"maxSize": 5242880,
				"protected": true
			}
		}`), edit_attachment); err != nil {
			return err
		}
		collection.Schema.AddField(edit_attachment)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("5hkmlmpqsznb4t4")
		if err != nil {
			return err
		}

		// update
		edit_attachment := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "9kqgwwui",
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
				"protected": false
			}
		}`), edit_attachment); err != nil {
			return err
		}
		collection.Schema.AddField(edit_attachment)

		return dao.SaveCollection(collection)
	})
}
