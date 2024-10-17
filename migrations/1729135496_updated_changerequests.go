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
		collection.Schema.RemoveField("z69rl5tf")

		// remove
		collection.Schema.RemoveField("vrjzon0s")

		// add
		new_summary := &schema.SchemaField{}
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
		}`), new_summary); err != nil {
			return err
		}
		collection.Schema.AddField(new_summary)

		// add
		new_listOfServices := &schema.SchemaField{}
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
		}`), new_listOfServices); err != nil {
			return err
		}
		collection.Schema.AddField(new_listOfServices)

		// add
		new_awarenessToBeMade := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "xap6irdp",
			"name": "awarenessToBeMade",
			"type": "select",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"Yes",
					"No",
					"Internal",
					"External",
					"Media"
				]
			}
		}`), new_awarenessToBeMade); err != nil {
			return err
		}
		collection.Schema.AddField(new_awarenessToBeMade)

		// add
		new_isClosed := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "g3ia97dz",
			"name": "isClosed",
			"type": "bool",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), new_isClosed); err != nil {
			return err
		}
		collection.Schema.AddField(new_isClosed)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("67689ye74w7y9pa")
		if err != nil {
			return err
		}

		// add
		del_date := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "z69rl5tf",
			"name": "date",
			"type": "date",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), del_date); err != nil {
			return err
		}
		collection.Schema.AddField(del_date)

		// add
		del_duration := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "vrjzon0s",
			"name": "duration",
			"type": "number",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"noDecimal": true
			}
		}`), del_duration); err != nil {
			return err
		}
		collection.Schema.AddField(del_duration)

		// remove
		collection.Schema.RemoveField("bfimixc7")

		// remove
		collection.Schema.RemoveField("ajqdlu4t")

		// remove
		collection.Schema.RemoveField("xap6irdp")

		// remove
		collection.Schema.RemoveField("g3ia97dz")

		return dao.SaveCollection(collection)
	})
}
