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

		// update
		edit_requestee := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "uhhhh4fb",
			"name": "requestee",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "85gg9za10dpmgdy",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), edit_requestee); err != nil {
			return err
		}
		collection.Schema.AddField(edit_requestee)

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

		// update
		edit_changeTeamId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "vjfh01kz",
			"name": "changeTeamId",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "yv3cl86s3yci1co",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), edit_changeTeamId); err != nil {
			return err
		}
		collection.Schema.AddField(edit_changeTeamId)

		// update
		edit_summary := &schema.SchemaField{}
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
		}`), edit_summary); err != nil {
			return err
		}
		collection.Schema.AddField(edit_summary)

		// update
		edit_scopeOfWork := &schema.SchemaField{}
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
		}`), edit_scopeOfWork); err != nil {
			return err
		}
		collection.Schema.AddField(edit_scopeOfWork)

		// update
		edit_risksAndMitigations := &schema.SchemaField{}
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
		}`), edit_risksAndMitigations); err != nil {
			return err
		}
		collection.Schema.AddField(edit_risksAndMitigations)

		// update
		edit_rollbackProcedures := &schema.SchemaField{}
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
		}`), edit_rollbackProcedures); err != nil {
			return err
		}
		collection.Schema.AddField(edit_rollbackProcedures)

		// update
		edit_listOfServices := &schema.SchemaField{}
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
		}`), edit_listOfServices); err != nil {
			return err
		}
		collection.Schema.AddField(edit_listOfServices)

		// update
		edit_awarenessToBeMade := &schema.SchemaField{}
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
		}`), edit_awarenessToBeMade); err != nil {
			return err
		}
		collection.Schema.AddField(edit_awarenessToBeMade)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("67689ye74w7y9pa")
		if err != nil {
			return err
		}

		// update
		edit_requestee := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "uhhhh4fb",
			"name": "requestee",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "85gg9za10dpmgdy",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), edit_requestee); err != nil {
			return err
		}
		collection.Schema.AddField(edit_requestee)

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

		// update
		edit_changeTeamId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "vjfh01kz",
			"name": "changeTeamId",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "yv3cl86s3yci1co",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), edit_changeTeamId); err != nil {
			return err
		}
		collection.Schema.AddField(edit_changeTeamId)

		// update
		edit_summary := &schema.SchemaField{}
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
		}`), edit_summary); err != nil {
			return err
		}
		collection.Schema.AddField(edit_summary)

		// update
		edit_scopeOfWork := &schema.SchemaField{}
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
		}`), edit_scopeOfWork); err != nil {
			return err
		}
		collection.Schema.AddField(edit_scopeOfWork)

		// update
		edit_risksAndMitigations := &schema.SchemaField{}
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
		}`), edit_risksAndMitigations); err != nil {
			return err
		}
		collection.Schema.AddField(edit_risksAndMitigations)

		// update
		edit_rollbackProcedures := &schema.SchemaField{}
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
		}`), edit_rollbackProcedures); err != nil {
			return err
		}
		collection.Schema.AddField(edit_rollbackProcedures)

		// update
		edit_listOfServices := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "aq7d2ohj",
			"name": "listOfServices",
			"type": "editor",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"convertUrls": false
			}
		}`), edit_listOfServices); err != nil {
			return err
		}
		collection.Schema.AddField(edit_listOfServices)

		// update
		edit_awarenessToBeMade := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "crdwu321",
			"name": "awarenessToBeMade",
			"type": "select",
			"required": true,
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
		}`), edit_awarenessToBeMade); err != nil {
			return err
		}
		collection.Schema.AddField(edit_awarenessToBeMade)

		return dao.SaveCollection(collection)
	})
}
