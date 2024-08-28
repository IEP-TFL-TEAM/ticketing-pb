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

		collection, err := dao.FindCollectionByNameOrId("v0i45iaci8zjm58")
		if err != nil {
			return err
		}

		// add
		new_faultTypeId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "vm7jau3j",
			"name": "faultTypeId",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "r8sn8obg2gksxgd",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), new_faultTypeId); err != nil {
			return err
		}
		collection.Schema.AddField(new_faultTypeId)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("v0i45iaci8zjm58")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("vm7jau3j")

		return dao.SaveCollection(collection)
	})
}
