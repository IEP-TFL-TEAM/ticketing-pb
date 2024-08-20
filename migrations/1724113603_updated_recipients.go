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

		collection, err := dao.FindCollectionByNameOrId("c607t2ux8dbejkb")
		if err != nil {
			return err
		}

		// add
		new_isVerified := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "5j1lbni9",
			"name": "isVerified",
			"type": "bool",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), new_isVerified); err != nil {
			return err
		}
		collection.Schema.AddField(new_isVerified)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("c607t2ux8dbejkb")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("5j1lbni9")

		return dao.SaveCollection(collection)
	})
}
