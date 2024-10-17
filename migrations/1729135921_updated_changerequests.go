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
		collection.Schema.RemoveField("bfimixc7")

		// remove
		collection.Schema.RemoveField("ajqdlu4t")

		// add
		new_summary := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "obm3pfks",
			"name": "summary",
			"type": "editor",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"convertUrls": false
			}
		}`), new_summary); err != nil {
			return err
		}
		collection.Schema.AddField(new_summary)

		// add
		new_listOfServices := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "efzoddxb",
			"name": "listOfServices",
			"type": "editor",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"convertUrls": false
			}
		}`), new_listOfServices); err != nil {
			return err
		}
		collection.Schema.AddField(new_listOfServices)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("67689ye74w7y9pa")
		if err != nil {
			return err
		}

		// add
		del_summary := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "bfimixc7",
			"name": "summary",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), del_summary); err != nil {
			return err
		}
		collection.Schema.AddField(del_summary)

		// add
		del_listOfServices := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "ajqdlu4t",
			"name": "listOfServices",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), del_listOfServices); err != nil {
			return err
		}
		collection.Schema.AddField(del_listOfServices)

		// remove
		collection.Schema.RemoveField("obm3pfks")

		// remove
		collection.Schema.RemoveField("efzoddxb")

		return dao.SaveCollection(collection)
	})
}
