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

		// remove
		collection.Schema.RemoveField("vnuf0x1r")

		// remove
		collection.Schema.RemoveField("foptlmnf")

		// remove
		collection.Schema.RemoveField("bofv0ozo")

		// remove
		collection.Schema.RemoveField("jzbhqda1")

		// add
		new_categoryLevelId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "8gfahr20",
			"name": "categoryLevelId",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "t0djc6rrp2vtymy",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), new_categoryLevelId); err != nil {
			return err
		}
		collection.Schema.AddField(new_categoryLevelId)

		// add
		new_teamEquipmentId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "uct1dcix",
			"name": "teamEquipmentId",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "4fgpahxwll5pzuj",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), new_teamEquipmentId); err != nil {
			return err
		}
		collection.Schema.AddField(new_teamEquipmentId)

		// add
		new_regionId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "qmffhsfk",
			"name": "regionId",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "vvnj2k58yney9nn",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), new_regionId); err != nil {
			return err
		}
		collection.Schema.AddField(new_regionId)

		// add
		new_areaId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "hipbxxo3",
			"name": "areaId",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "63z0ygwtpjs28a4",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), new_areaId); err != nil {
			return err
		}
		collection.Schema.AddField(new_areaId)

		// add
		new_siteId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "dygdueds",
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
		new_faultTypeId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "adsoetyj",
			"name": "faultTypeId",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "r8sn8obg2gksxgd",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), new_faultTypeId); err != nil {
			return err
		}
		collection.Schema.AddField(new_faultTypeId)

		// update
		edit_status := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "b5p1otrz",
			"name": "status",
			"type": "select",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"PENDING",
					"APPROVED",
					"CLOSED"
				]
			}
		}`), edit_status); err != nil {
			return err
		}
		collection.Schema.AddField(edit_status)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("caza2oj0rtxw1gu")
		if err != nil {
			return err
		}

		// add
		del_staffId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "vnuf0x1r",
			"name": "staffId",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "vqywtenuwa86j8x",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), del_staffId); err != nil {
			return err
		}
		collection.Schema.AddField(del_staffId)

		// add
		del_severity := &schema.SchemaField{}
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
					"MINOR",
					"MAJOR",
					"CRITICAL"
				]
			}
		}`), del_severity); err != nil {
			return err
		}
		collection.Schema.AddField(del_severity)

		// add
		del_location := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "bofv0ozo",
			"name": "location",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), del_location); err != nil {
			return err
		}
		collection.Schema.AddField(del_location)

		// add
		del_locationId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "jzbhqda1",
			"name": "locationId",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "blal3gzv2ctzdg8",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), del_locationId); err != nil {
			return err
		}
		collection.Schema.AddField(del_locationId)

		// remove
		collection.Schema.RemoveField("8gfahr20")

		// remove
		collection.Schema.RemoveField("uct1dcix")

		// remove
		collection.Schema.RemoveField("qmffhsfk")

		// remove
		collection.Schema.RemoveField("hipbxxo3")

		// remove
		collection.Schema.RemoveField("dygdueds")

		// remove
		collection.Schema.RemoveField("adsoetyj")

		// update
		edit_status := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "b5p1otrz",
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
		}`), edit_status); err != nil {
			return err
		}
		collection.Schema.AddField(edit_status)

		return dao.SaveCollection(collection)
	})
}
