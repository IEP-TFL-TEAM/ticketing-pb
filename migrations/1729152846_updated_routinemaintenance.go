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

		// update
		edit_isClosed := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "rb9pc45j",
			"name": "isClosed",
			"type": "bool",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), edit_isClosed); err != nil {
			return err
		}
		collection.Schema.AddField(edit_isClosed)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("w0xy9jxjhwjgnz3")
		if err != nil {
			return err
		}

		// update
		edit_isClosed := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "rb9pc45j",
			"name": "closed",
			"type": "bool",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), edit_isClosed); err != nil {
			return err
		}
		collection.Schema.AddField(edit_isClosed)

		return dao.SaveCollection(collection)
	})
}
