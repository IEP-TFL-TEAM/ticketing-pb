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
		new_description := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "uvhtzykh",
			"name": "description",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_description); err != nil {
			return err
		}
		collection.Schema.AddField(new_description)

		// add
		new_reportedBy := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "d0oj3dri",
			"name": "reportedBy",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "_pb_users_auth_",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), new_reportedBy); err != nil {
			return err
		}
		collection.Schema.AddField(new_reportedBy)

		// add
		new_teamId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "f0ojadtu",
			"name": "teamId",
			"type": "relation",
			"required": false,
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
		new_status := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "8wqtkbba",
			"name": "status",
			"type": "select",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"OPEN",
					"COMPLETE",
					"IN PROGRESS",
					"CLOSED",
					"PENDING",
					"RESOLVED"
				]
			}
		}`), new_status); err != nil {
			return err
		}
		collection.Schema.AddField(new_status)

		// add
		new_severity := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "foptlmnf",
			"name": "severity",
			"type": "select",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"LOW",
					"MEDIUM",
					"HIGH",
					"SEVERE"
				]
			}
		}`), new_severity); err != nil {
			return err
		}
		collection.Schema.AddField(new_severity)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("caza2oj0rtxw1gu")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("uvhtzykh")

		// remove
		collection.Schema.RemoveField("d0oj3dri")

		// remove
		collection.Schema.RemoveField("f0ojadtu")

		// remove
		collection.Schema.RemoveField("8wqtkbba")

		// remove
		collection.Schema.RemoveField("foptlmnf")

		return dao.SaveCollection(collection)
	})
}
