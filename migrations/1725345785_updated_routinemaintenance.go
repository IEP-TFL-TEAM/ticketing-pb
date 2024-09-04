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
		collection.Schema.RemoveField("gocwlo66")

		// add
		new_requestee := &schema.SchemaField{}
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
		}`), new_requestee); err != nil {
			return err
		}
		collection.Schema.AddField(new_requestee)

		// add
		new_date := &schema.SchemaField{}
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
		}`), new_date); err != nil {
			return err
		}
		collection.Schema.AddField(new_date)

		// add
		new_title := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "xx4nl8ld",
			"name": "title",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_title); err != nil {
			return err
		}
		collection.Schema.AddField(new_title)

		// add
		new_objective := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "sdm0ora9",
			"name": "objective",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_objective); err != nil {
			return err
		}
		collection.Schema.AddField(new_objective)

		// add
		new_riskAversion := &schema.SchemaField{}
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
		}`), new_riskAversion); err != nil {
			return err
		}
		collection.Schema.AddField(new_riskAversion)

		// add
		new_scopeOfWork := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "1tnfqxdw",
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
		new_siteId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "0rl6ali6",
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

		// add
		new_startDate := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "xkyxcsf0",
			"name": "startDate",
			"type": "date",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), new_startDate); err != nil {
			return err
		}
		collection.Schema.AddField(new_startDate)

		// add
		new_endDate := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "7aa7g2by",
			"name": "endDate",
			"type": "date",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), new_endDate); err != nil {
			return err
		}
		collection.Schema.AddField(new_endDate)

		// add
		new_serviceImpact := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "oztlvyel",
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
		new_duration := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "jcmaqvg5",
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
		}`), new_duration); err != nil {
			return err
		}
		collection.Schema.AddField(new_duration)

		// add
		new_teamIds := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "wwitwnyi",
			"name": "teamIds",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "rybge8g8z4jgrwi",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": null,
				"displayFields": null
			}
		}`), new_teamIds); err != nil {
			return err
		}
		collection.Schema.AddField(new_teamIds)

		// add
		new_maintenanceTeamId := &schema.SchemaField{}
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
		}`), new_maintenanceTeamId); err != nil {
			return err
		}
		collection.Schema.AddField(new_maintenanceTeamId)

		// add
		new_listOfServices := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "ltbmqhlb",
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

		// add
		new_awarenessToBeMade := &schema.SchemaField{}
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
		}`), new_awarenessToBeMade); err != nil {
			return err
		}
		collection.Schema.AddField(new_awarenessToBeMade)

		// add
		new_maintenanceRisks := &schema.SchemaField{}
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
		}`), new_maintenanceRisks); err != nil {
			return err
		}
		collection.Schema.AddField(new_maintenanceRisks)

		// add
		new_maintenancePreChecklist := &schema.SchemaField{}
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
		}`), new_maintenancePreChecklist); err != nil {
			return err
		}
		collection.Schema.AddField(new_maintenancePreChecklist)

		// add
		new_maintenanceOutcome := &schema.SchemaField{}
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
		}`), new_maintenanceOutcome); err != nil {
			return err
		}
		collection.Schema.AddField(new_maintenanceOutcome)

		// add
		new_attachment := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "ec2bfeml",
			"name": "attachment",
			"type": "file",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"mimeTypes": [
					"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
					"image/png",
					"image/vnd.mozilla.apng",
					"image/jpeg",
					"application/msword",
					"application/pdf"
				],
				"thumbs": [],
				"maxSelect": 1,
				"maxSize": 5242880,
				"protected": true
			}
		}`), new_attachment); err != nil {
			return err
		}
		collection.Schema.AddField(new_attachment)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("w0xy9jxjhwjgnz3")
		if err != nil {
			return err
		}

		// add
		del_requestee := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "gocwlo66",
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

		// remove
		collection.Schema.RemoveField("jchr9asn")

		// remove
		collection.Schema.RemoveField("g98fbkng")

		// remove
		collection.Schema.RemoveField("xx4nl8ld")

		// remove
		collection.Schema.RemoveField("sdm0ora9")

		// remove
		collection.Schema.RemoveField("3zuvuhrf")

		// remove
		collection.Schema.RemoveField("1tnfqxdw")

		// remove
		collection.Schema.RemoveField("0rl6ali6")

		// remove
		collection.Schema.RemoveField("xkyxcsf0")

		// remove
		collection.Schema.RemoveField("7aa7g2by")

		// remove
		collection.Schema.RemoveField("oztlvyel")

		// remove
		collection.Schema.RemoveField("jcmaqvg5")

		// remove
		collection.Schema.RemoveField("wwitwnyi")

		// remove
		collection.Schema.RemoveField("mprpgron")

		// remove
		collection.Schema.RemoveField("ltbmqhlb")

		// remove
		collection.Schema.RemoveField("zpupxxeq")

		// remove
		collection.Schema.RemoveField("wjhc2t9q")

		// remove
		collection.Schema.RemoveField("jdvrqcam")

		// remove
		collection.Schema.RemoveField("awauzk8a")

		// remove
		collection.Schema.RemoveField("ec2bfeml")

		return dao.SaveCollection(collection)
	})
}
