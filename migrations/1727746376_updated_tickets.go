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

		// add
		new_incidentEnd := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "innrrgmg",
			"name": "incidentEnd",
			"type": "date",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), new_incidentEnd); err != nil {
			return err
		}
		collection.Schema.AddField(new_incidentEnd)

		// update
		edit_slaStatus := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "dv5ibxy3",
			"name": "slaStatus",
			"type": "select",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"SLA Met",
					"SLA Exceeded"
				]
			}
		}`), edit_slaStatus); err != nil {
			return err
		}
		collection.Schema.AddField(edit_slaStatus)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("caza2oj0rtxw1gu")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("innrrgmg")

		// update
		edit_slaStatus := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "dv5ibxy3",
			"name": "slaStatus",
			"type": "select",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"Within SLA",
					"SLA Exceeded"
				]
			}
		}`), edit_slaStatus); err != nil {
			return err
		}
		collection.Schema.AddField(edit_slaStatus)

		return dao.SaveCollection(collection)
	})
}
