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

		collection, err := dao.FindCollectionByNameOrId("6v2l2d2974asxtv")
		if err != nil {
			return err
		}

		// update
		edit_isLead := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "p0axpg9c",
			"name": "isLead",
			"type": "bool",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), edit_isLead); err != nil {
			return err
		}
		collection.Schema.AddField(edit_isLead)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("6v2l2d2974asxtv")
		if err != nil {
			return err
		}

		// update
		edit_isLead := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "p0axpg9c",
			"name": "isLead",
			"type": "bool",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), edit_isLead); err != nil {
			return err
		}
		collection.Schema.AddField(edit_isLead)

		return dao.SaveCollection(collection)
	})
}
