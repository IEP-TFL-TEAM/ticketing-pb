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
		collection.Schema.RemoveField("cnuym4up")

		// remove
		collection.Schema.RemoveField("exasvciw")

		// remove
		collection.Schema.RemoveField("qritmvwp")

		// remove
		collection.Schema.RemoveField("j78ngjye")

		// add
		new_summary := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "q66fevaq",
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
		new_scopeOfWork := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "9rgko3ff",
			"name": "scopeOfWork",
			"type": "editor",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"convertUrls": false
			}
		}`), new_scopeOfWork); err != nil {
			return err
		}
		collection.Schema.AddField(new_scopeOfWork)

		// add
		new_risksAndMitigations := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "zl81le2p",
			"name": "risksAndMitigations",
			"type": "editor",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"convertUrls": false
			}
		}`), new_risksAndMitigations); err != nil {
			return err
		}
		collection.Schema.AddField(new_risksAndMitigations)

		// add
		new_rollbackProcedures := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "p8q8ntlp",
			"name": "rollbackProcedures",
			"type": "editor",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"convertUrls": false
			}
		}`), new_rollbackProcedures); err != nil {
			return err
		}
		collection.Schema.AddField(new_rollbackProcedures)

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
			"id": "cnuym4up",
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
		del_scopeOfWork := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "exasvciw",
			"name": "scopeOfWork",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), del_scopeOfWork); err != nil {
			return err
		}
		collection.Schema.AddField(del_scopeOfWork)

		// add
		del_risksAndMitigations := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "qritmvwp",
			"name": "risksAndMitigations",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), del_risksAndMitigations); err != nil {
			return err
		}
		collection.Schema.AddField(del_risksAndMitigations)

		// add
		del_rollbackProcedures := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "j78ngjye",
			"name": "rollbackProcedures",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), del_rollbackProcedures); err != nil {
			return err
		}
		collection.Schema.AddField(del_rollbackProcedures)

		// remove
		collection.Schema.RemoveField("q66fevaq")

		// remove
		collection.Schema.RemoveField("9rgko3ff")

		// remove
		collection.Schema.RemoveField("zl81le2p")

		// remove
		collection.Schema.RemoveField("p8q8ntlp")

		return dao.SaveCollection(collection)
	})
}
