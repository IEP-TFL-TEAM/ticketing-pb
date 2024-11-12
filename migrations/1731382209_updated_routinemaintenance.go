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
		edit_duration := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "jcmaqvg5",
			"name": "duration",
			"type": "number",
			"required": false,
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

		collection, err := dao.FindCollectionByNameOrId("w0xy9jxjhwjgnz3")
		if err != nil {
			return err
		}

		// update
		edit_duration := &schema.SchemaField{}
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
		}`), edit_duration); err != nil {
			return err
		}
		collection.Schema.AddField(edit_duration)

		return dao.SaveCollection(collection)
	})
}
