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

		collection, err := dao.FindCollectionByNameOrId("67689ye74w7y9pa")
		if err != nil {
			return err
		}

		// update
		edit_endDate := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "cotbbonr",
			"name": "endDate",
			"type": "date",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), edit_endDate); err != nil {
			return err
		}
		collection.Schema.AddField(edit_endDate)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("67689ye74w7y9pa")
		if err != nil {
			return err
		}

		// update
		edit_endDate := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "cotbbonr",
			"name": "endDate",
			"type": "date",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), edit_endDate); err != nil {
			return err
		}
		collection.Schema.AddField(edit_endDate)

		return dao.SaveCollection(collection)
	})
}
