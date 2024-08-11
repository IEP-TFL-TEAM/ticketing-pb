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

		collection, err := dao.FindCollectionByNameOrId("ssdqmjdubsma48q")
		if err != nil {
			return err
		}

		collection.Name = "equipment"

		// remove
		collection.Schema.RemoveField("xnpwnguu")

		// add
		new_siteId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "jjqfth5s",
			"name": "siteId",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "ghtuohic73e7d35",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), new_siteId); err != nil {
			return err
		}
		collection.Schema.AddField(new_siteId)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("ssdqmjdubsma48q")
		if err != nil {
			return err
		}

		collection.Name = "materials"

		// add
		del_quantity := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "xnpwnguu",
			"name": "quantity",
			"type": "number",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"noDecimal": true
			}
		}`), del_quantity); err != nil {
			return err
		}
		collection.Schema.AddField(del_quantity)

		// remove
		collection.Schema.RemoveField("jjqfth5s")

		return dao.SaveCollection(collection)
	})
}
