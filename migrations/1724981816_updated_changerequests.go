package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("67689ye74w7y9pa")
		if err != nil {
			return err
		}

		collection.CreateRule = types.Pointer("@request.auth.id != \"\" && @request.auth.role = \"admin\"")

		collection.UpdateRule = types.Pointer("@request.auth.id != \"\" && @request.auth.role = \"admin\"")

		// remove
		collection.Schema.RemoveField("j0clm1si")

		// remove
		collection.Schema.RemoveField("jgv8tfbp")

		// remove
		collection.Schema.RemoveField("lmqvgiiz")

		// remove
		collection.Schema.RemoveField("8sdh8d2x")

		// add
		new_requestedBy := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "uhhhh4fb",
			"name": "requestedBy",
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
		}`), new_requestedBy); err != nil {
			return err
		}
		collection.Schema.AddField(new_requestedBy)

		// add
		new_date := &schema.SchemaField{}
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
		}`), new_date); err != nil {
			return err
		}
		collection.Schema.AddField(new_date)

		// add
		new_serviceImpact := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "j8cox1ub",
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
		new_teamId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "c51dwpud",
			"name": "teamId",
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
		}`), new_teamId); err != nil {
			return err
		}
		collection.Schema.AddField(new_teamId)

		// add
		new_risksAndMitigations := &schema.SchemaField{}
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
		}`), new_risksAndMitigations); err != nil {
			return err
		}
		collection.Schema.AddField(new_risksAndMitigations)

		// add
		new_rollbackProcedures := &schema.SchemaField{}
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
		}`), new_rollbackProcedures); err != nil {
			return err
		}
		collection.Schema.AddField(new_rollbackProcedures)

		// update
		edit_duration := &schema.SchemaField{}
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
				"noDecimal": false
			}
		}`), edit_duration); err != nil {
			return err
		}
		collection.Schema.AddField(edit_duration)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("67689ye74w7y9pa")
		if err != nil {
			return err
		}

		collection.CreateRule = types.Pointer("@request.auth.id != \"\"")

		collection.UpdateRule = types.Pointer("@request.auth.id != \"\"")

		// add
		del_involvedTechGroups := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "j0clm1si",
			"name": "involvedTechGroups",
			"type": "select",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 16,
				"values": [
					"CAN CE",
					"CAN WE",
					"CAN NO",
					"Access AOM CE",
					"Access AOM WE",
					"Access AOM NO",
					"Transport",
					"FVDN",
					"ISP",
					"Telepower",
					"Transport",
					"Wireless",
					"Enterprise Solutions & Projects",
					"GIT",
					"Contractor",
					"Vendor"
				]
			}
		}`), del_involvedTechGroups); err != nil {
			return err
		}
		collection.Schema.AddField(del_involvedTechGroups)

		// add
		del_date := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "jgv8tfbp",
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
		del_requestee := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "lmqvgiiz",
			"name": "requestee",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), del_requestee); err != nil {
			return err
		}
		collection.Schema.AddField(del_requestee)

		// add
		del_serviceImpact := &schema.SchemaField{}
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
		}`), del_serviceImpact); err != nil {
			return err
		}
		collection.Schema.AddField(del_serviceImpact)

		// remove
		collection.Schema.RemoveField("uhhhh4fb")

		// remove
		collection.Schema.RemoveField("z69rl5tf")

		// remove
		collection.Schema.RemoveField("j8cox1ub")

		// remove
		collection.Schema.RemoveField("c51dwpud")

		// remove
		collection.Schema.RemoveField("qritmvwp")

		// remove
		collection.Schema.RemoveField("j78ngjye")

		// update
		edit_duration := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "vrjzon0s",
			"name": "duration",
			"type": "number",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"noDecimal": false
			}
		}`), edit_duration); err != nil {
			return err
		}
		collection.Schema.AddField(edit_duration)

		return dao.SaveCollection(collection)
	})
}
