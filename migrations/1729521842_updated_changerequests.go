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

		collection, err := dao.FindCollectionByNameOrId("67689ye74w7y9pa")
		if err != nil {
			return err
		}

		// add
		new_closingRemarks := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "34cuwl3h",
			"name": "closingRemarks",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_closingRemarks); err != nil {
			return err
		}
		collection.Schema.AddField(new_closingRemarks)

		// add
		new_closingAttachment := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "rixlzh0j",
			"name": "closingAttachment",
			"type": "file",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"mimeTypes": [],
				"thumbs": [],
				"maxSelect": 1,
				"maxSize": 5242880,
				"protected": true
			}
		}`), new_closingAttachment); err != nil {
			return err
		}
		collection.Schema.AddField(new_closingAttachment)

		// update
		edit_attachment := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "gv5cz7tu",
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
				"protected": true
			}
		}`), edit_attachment); err != nil {
			return err
		}
		collection.Schema.AddField(edit_attachment)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("67689ye74w7y9pa")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("34cuwl3h")

		// remove
		collection.Schema.RemoveField("rixlzh0j")

		// update
		edit_attachment := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "gv5cz7tu",
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
					"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
					"application/msword",
					"application/pdf"
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
