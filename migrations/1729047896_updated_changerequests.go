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
		collection.Schema.RemoveField("q66fevaq")

		// remove
		collection.Schema.RemoveField("9rgko3ff")

		// remove
		collection.Schema.RemoveField("zl81le2p")

		// remove
		collection.Schema.RemoveField("p8q8ntlp")

		// remove
		collection.Schema.RemoveField("aq7d2ohj")

		// remove
		collection.Schema.RemoveField("crdwu321")

		// update
		edit_date := &schema.SchemaField{}
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
		}`), edit_date); err != nil {
			return err
		}
		collection.Schema.AddField(edit_date)

		// update
		edit_startDate := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "v0w7jtth",
			"name": "startDate",
			"type": "date",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), edit_startDate); err != nil {
			return err
		}
		collection.Schema.AddField(edit_startDate)

		// update
		edit_endDate := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "cotbbonr",
			"name": "endDate",
			"type": "date",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), edit_endDate); err != nil {
			return err
		}
		collection.Schema.AddField(edit_endDate)

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
			"id": "q66fevaq",
			"name": "summary",
			"type": "editor",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"convertUrls": false
			}
		}`), del_summary); err != nil {
			return err
		}
		collection.Schema.AddField(del_summary)

		// add
		del_scopeOfWork := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "9rgko3ff",
			"name": "scopeOfWork",
			"type": "editor",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"convertUrls": false
			}
		}`), del_scopeOfWork); err != nil {
			return err
		}
		collection.Schema.AddField(del_scopeOfWork)

		// add
		del_risksAndMitigations := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "zl81le2p",
			"name": "risksAndMitigations",
			"type": "editor",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"convertUrls": false
			}
		}`), del_risksAndMitigations); err != nil {
			return err
		}
		collection.Schema.AddField(del_risksAndMitigations)

		// add
		del_rollbackProcedures := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "p8q8ntlp",
			"name": "rollbackProcedures",
			"type": "editor",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"convertUrls": false
			}
		}`), del_rollbackProcedures); err != nil {
			return err
		}
		collection.Schema.AddField(del_rollbackProcedures)

		// add
		del_listOfServices := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "aq7d2ohj",
			"name": "listOfServices",
			"type": "editor",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"convertUrls": false
			}
		}`), del_listOfServices); err != nil {
			return err
		}
		collection.Schema.AddField(del_listOfServices)

		// add
		del_awarenessToBeMade := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "crdwu321",
			"name": "awarenessToBeMade",
			"type": "select",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"No",
					"Internal",
					"External",
					"Media"
				]
			}
		}`), del_awarenessToBeMade); err != nil {
			return err
		}
		collection.Schema.AddField(del_awarenessToBeMade)

		// update
		edit_date := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "z69rl5tf",
			"name": "date",
			"type": "date",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), edit_date); err != nil {
			return err
		}
		collection.Schema.AddField(edit_date)

		// update
		edit_startDate := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "v0w7jtth",
			"name": "startDate",
			"type": "date",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), edit_startDate); err != nil {
			return err
		}
		collection.Schema.AddField(edit_startDate)

		// update
		edit_endDate := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "cotbbonr",
			"name": "endDate",
			"type": "date",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), edit_endDate); err != nil {
			return err
		}
		collection.Schema.AddField(edit_endDate)

		return dao.SaveCollection(collection)
	})
}
