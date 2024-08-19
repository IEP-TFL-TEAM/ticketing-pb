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

		// remove
		collection.Schema.RemoveField("rnduiv6s")

		// add
		new_requestee := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "lmqvgiiz",
			"name": "requestee",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_requestee); err != nil {
			return err
		}
		collection.Schema.AddField(new_requestee)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("67689ye74w7y9pa")
		if err != nil {
			return err
		}

		// add
		del_requestee := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "rnduiv6s",
			"name": "requestee",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "_pb_users_auth_",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), del_requestee); err != nil {
			return err
		}
		collection.Schema.AddField(del_requestee)

		// remove
		collection.Schema.RemoveField("lmqvgiiz")

		return dao.SaveCollection(collection)
	})
}
