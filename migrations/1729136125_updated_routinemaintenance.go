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

		collection, err := dao.FindCollectionByNameOrId("w0xy9jxjhwjgnz3")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("g98fbkng")

		// remove
		collection.Schema.RemoveField("3zuvuhrf")

		// remove
		collection.Schema.RemoveField("zpupxxeq")

		// remove
		collection.Schema.RemoveField("wjhc2t9q")

		// remove
		collection.Schema.RemoveField("jdvrqcam")

		// remove
		collection.Schema.RemoveField("awauzk8a")

		// add
		new_closed := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "rb9pc45j",
			"name": "closed",
			"type": "bool",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), new_closed); err != nil {
			return err
		}
		collection.Schema.AddField(new_closed)

		// update
		edit_requestee := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "jchr9asn",
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
		edit_maintenanceTeamId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "mprpgron",
			"name": "maintenanceTeamId",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "16udx7u72g9xwti",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), edit_maintenanceTeamId); err != nil {
			return err
		}
		collection.Schema.AddField(edit_maintenanceTeamId)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("w0xy9jxjhwjgnz3")
		if err != nil {
			return err
		}

		// add
		del_date := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "g98fbkng",
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
		del_riskAversion := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "3zuvuhrf",
			"name": "riskAversion",
			"type": "editor",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"convertUrls": false
			}
		}`), del_riskAversion); err != nil {
			return err
		}
		collection.Schema.AddField(del_riskAversion)

		// add
		del_awarenessToBeMade := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "zpupxxeq",
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
		}`), del_awarenessToBeMade); err != nil {
			return err
		}
		collection.Schema.AddField(del_awarenessToBeMade)

		// add
		del_maintenanceRisks := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "wjhc2t9q",
			"name": "maintenanceRisks",
			"type": "editor",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"convertUrls": false
			}
		}`), del_maintenanceRisks); err != nil {
			return err
		}
		collection.Schema.AddField(del_maintenanceRisks)

		// add
		del_maintenancePreChecklist := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "jdvrqcam",
			"name": "maintenancePreChecklist",
			"type": "editor",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"convertUrls": false
			}
		}`), del_maintenancePreChecklist); err != nil {
			return err
		}
		collection.Schema.AddField(del_maintenancePreChecklist)

		// add
		del_maintenanceOutcome := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "awauzk8a",
			"name": "maintenanceOutcome",
			"type": "editor",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"convertUrls": false
			}
		}`), del_maintenanceOutcome); err != nil {
			return err
		}
		collection.Schema.AddField(del_maintenanceOutcome)

		// remove
		collection.Schema.RemoveField("rb9pc45j")

		// update
		edit_requestee := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "jchr9asn",
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
		edit_maintenanceTeamId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "mprpgron",
			"name": "maintenanceTeamId",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "16udx7u72g9xwti",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), edit_maintenanceTeamId); err != nil {
			return err
		}
		collection.Schema.AddField(edit_maintenanceTeamId)

		return dao.SaveCollection(collection)
	})
}
