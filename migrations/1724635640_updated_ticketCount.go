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

		collection, err := dao.FindCollectionByNameOrId("l5qc8lbi3y6q2s7")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("@request.auth.id != \"\"")

		collection.ViewRule = types.Pointer("@request.auth.id != \"\"")

		// remove
		collection.Schema.RemoveField("awiv4hv0")

		// add
		new_totalItems := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "hnbwiyji",
			"name": "totalItems",
			"type": "number",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"noDecimal": false
			}
		}`), new_totalItems); err != nil {
			return err
		}
		collection.Schema.AddField(new_totalItems)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("l5qc8lbi3y6q2s7")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("@request.auth.collectionName = \"users\"")

		collection.ViewRule = types.Pointer("@request.auth.collectionName = \"users\"")

		// add
		del_totalItems := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "awiv4hv0",
			"name": "totalItems",
			"type": "number",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"noDecimal": false
			}
		}`), del_totalItems); err != nil {
			return err
		}
		collection.Schema.AddField(del_totalItems)

		// remove
		collection.Schema.RemoveField("hnbwiyji")

		return dao.SaveCollection(collection)
	})
}
