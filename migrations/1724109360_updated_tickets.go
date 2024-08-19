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
		collection.Schema.RemoveField("rjiedfpj")

		// remove
		collection.Schema.RemoveField("vn3wd0vt")

		// add
		new_technicianId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "zktfi8or",
			"name": "technicianId",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "a3vplmoundlez9b",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), new_technicianId); err != nil {
			return err
		}
		collection.Schema.AddField(new_technicianId)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("caza2oj0rtxw1gu")
		if err != nil {
			return err
		}

		// add
		del_isVerified := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "rjiedfpj",
			"name": "isVerified",
			"type": "bool",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), del_isVerified); err != nil {
			return err
		}
		collection.Schema.AddField(del_isVerified)

		// add
		del_isSpecial := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "vn3wd0vt",
			"name": "isSpecial",
			"type": "bool",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), del_isSpecial); err != nil {
			return err
		}
		collection.Schema.AddField(del_isSpecial)

		// remove
		collection.Schema.RemoveField("zktfi8or")

		return dao.SaveCollection(collection)
	})
}
