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

		collection, err := dao.FindCollectionByNameOrId("blal3gzv2ctzdg8")
		if err != nil {
			return err
		}

		// update
		edit_latitude := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "arllkzxo",
			"name": "latitude",
			"type": "number",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"noDecimal": false
			}
		}`), edit_latitude); err != nil {
			return err
		}
		collection.Schema.AddField(edit_latitude)

		// update
		edit_longitude := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "kpt5atb0",
			"name": "longitude",
			"type": "number",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"noDecimal": false
			}
		}`), edit_longitude); err != nil {
			return err
		}
		collection.Schema.AddField(edit_longitude)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("blal3gzv2ctzdg8")
		if err != nil {
			return err
		}

		// update
		edit_latitude := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "arllkzxo",
			"name": "latitude",
			"type": "number",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"noDecimal": false
			}
		}`), edit_latitude); err != nil {
			return err
		}
		collection.Schema.AddField(edit_latitude)

		// update
		edit_longitude := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "kpt5atb0",
			"name": "longitude",
			"type": "number",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"noDecimal": false
			}
		}`), edit_longitude); err != nil {
			return err
		}
		collection.Schema.AddField(edit_longitude)

		return dao.SaveCollection(collection)
	})
}
