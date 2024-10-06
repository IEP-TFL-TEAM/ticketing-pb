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

		collection, err := dao.FindCollectionByNameOrId("5mec029rdzm42xd")
		if err != nil {
			return err
		}

		// update
		edit_attachment := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "x57kgvpt",
			"name": "attachment",
			"type": "file",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"mimeTypes": [
					"image/png",
					"image/vnd.mozilla.apng",
					"image/jpeg",
					"image/webp",
					"image/svg+xml",
					"image/gif"
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

		collection, err := dao.FindCollectionByNameOrId("5mec029rdzm42xd")
		if err != nil {
			return err
		}

		// update
		edit_attachment := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "x57kgvpt",
			"name": "attachment",
			"type": "file",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"mimeTypes": [
					"application/pdf",
					"image/png",
					"image/vnd.mozilla.apng",
					"image/jpeg",
					"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
					"application/msword"
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
	})
}