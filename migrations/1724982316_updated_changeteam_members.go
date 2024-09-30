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

		collection, err := dao.FindCollectionByNameOrId("6v2l2d2974asxtv")
		if err != nil {
			return err
		}

		// add
		new_changeTeamId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "qz2bq2qr",
			"name": "changeTeamId",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "yv3cl86s3yci1co",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), new_changeTeamId); err != nil {
			return err
		}
		collection.Schema.AddField(new_changeTeamId)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("6v2l2d2974asxtv")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("qz2bq2qr")

		return dao.SaveCollection(collection)
	})
}