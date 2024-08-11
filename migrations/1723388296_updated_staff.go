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

		collection, err := dao.FindCollectionByNameOrId("vqywtenuwa86j8x")
		if err != nil {
			return err
		}

		// add
		new_teamId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "g2t9g5d5",
			"name": "teamId",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "rybge8g8z4jgrwi",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), new_teamId); err != nil {
			return err
		}
		collection.Schema.AddField(new_teamId)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("vqywtenuwa86j8x")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("g2t9g5d5")

		return dao.SaveCollection(collection)
	})
}
