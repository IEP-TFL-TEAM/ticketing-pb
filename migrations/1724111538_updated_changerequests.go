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
		collection.Schema.RemoveField("zdv08arn")

		// remove
		collection.Schema.RemoveField("g34l17i5")

		// remove
		collection.Schema.RemoveField("nxbnpyik")

		// add
		new_serviceImpact := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "8sdh8d2x",
			"name": "serviceImpact",
			"type": "select",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"Yes",
					"No"
				]
			}
		}`), new_serviceImpact); err != nil {
			return err
		}
		collection.Schema.AddField(new_serviceImpact)

		// add
		new_siteId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "sjeypqhd",
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

		collection, err := dao.FindCollectionByNameOrId("67689ye74w7y9pa")
		if err != nil {
			return err
		}

		// add
		del_serviceImpact := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "zdv08arn",
			"name": "serviceImpact",
			"type": "bool",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), del_serviceImpact); err != nil {
			return err
		}
		collection.Schema.AddField(del_serviceImpact)

		// add
		del_taskSite := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "g34l17i5",
			"name": "taskSite",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), del_taskSite); err != nil {
			return err
		}
		collection.Schema.AddField(del_taskSite)

		// add
		del_changeTeam := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "nxbnpyik",
			"name": "changeTeam",
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
		}`), del_changeTeam); err != nil {
			return err
		}
		collection.Schema.AddField(del_changeTeam)

		// remove
		collection.Schema.RemoveField("8sdh8d2x")

		// remove
		collection.Schema.RemoveField("sjeypqhd")

		return dao.SaveCollection(collection)
	})
}
