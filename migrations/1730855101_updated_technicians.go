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

		collection, err := dao.FindCollectionByNameOrId("a3vplmoundlez9b")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("tb6typsh")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("a3vplmoundlez9b")
		if err != nil {
			return err
		}

		// add
		del_teamId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "tb6typsh",
			"name": "teamId",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "rybge8g8z4jgrwi",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), del_teamId); err != nil {
			return err
		}
		collection.Schema.AddField(del_teamId)

		return dao.SaveCollection(collection)
	})
}
