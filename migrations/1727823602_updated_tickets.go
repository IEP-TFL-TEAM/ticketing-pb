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

		// update
		edit_teamEquipmentIds := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "uct1dcix",
			"name": "teamEquipmentIds",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "4fgpahxwll5pzuj",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": null,
				"displayFields": null
			}
		}`), edit_teamEquipmentIds); err != nil {
			return err
		}
		collection.Schema.AddField(edit_teamEquipmentIds)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("caza2oj0rtxw1gu")
		if err != nil {
			return err
		}

		// update
		edit_teamEquipmentIds := &schema.SchemaField{}
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
				"maxSelect": null,
				"displayFields": null
			}
		}`), edit_teamEquipmentIds); err != nil {
			return err
		}
		collection.Schema.AddField(edit_teamEquipmentIds)

		return dao.SaveCollection(collection)
	})
}
