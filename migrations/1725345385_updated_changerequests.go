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

		collection, err := dao.FindCollectionByNameOrId("67689ye74w7y9pa")
		if err != nil {
			return err
		}

		collection.CreateRule = types.Pointer("@request.auth.id != \"\"")

		collection.UpdateRule = types.Pointer("@request.auth.id != \"\"")

		// update
		edit_duration := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "vrjzon0s",
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
		}`), edit_duration); err != nil {
			return err
		}
		collection.Schema.AddField(edit_duration)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("67689ye74w7y9pa")
		if err != nil {
			return err
		}

		collection.CreateRule = types.Pointer("@request.auth.id != \"\" && @request.auth.role = \"admin\"")

		collection.UpdateRule = types.Pointer("@request.auth.id != \"\" && @request.auth.role = \"admin\"")

		// update
		edit_duration := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "vrjzon0s",
			"name": "duration",
			"type": "number",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"noDecimal": false
			}
		}`), edit_duration); err != nil {
			return err
		}
		collection.Schema.AddField(edit_duration)

		return dao.SaveCollection(collection)
	})
}
