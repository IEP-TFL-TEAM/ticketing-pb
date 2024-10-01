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

		collection, err := dao.FindCollectionByNameOrId("caza2oj0rtxw1gu")
		if err != nil {
			return err
		}

		// update
		edit_causedBy := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "qxfuqxxw",
			"name": "causedBy",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "v0i45iaci8zjm58",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), edit_causedBy); err != nil {
			return err
		}
		collection.Schema.AddField(edit_causedBy)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("caza2oj0rtxw1gu")
		if err != nil {
			return err
		}

		// update
		edit_causedBy := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "qxfuqxxw",
			"name": "causedBy",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "v0i45iaci8zjm58",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), edit_causedBy); err != nil {
			return err
		}
		collection.Schema.AddField(edit_causedBy)

		return dao.SaveCollection(collection)
	})
}
