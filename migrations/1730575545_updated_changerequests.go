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
		edit_awarenessToBeMade := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "xap6irdp",
			"name": "awarenessToBeMade",
			"type": "select",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 5,
				"values": [
					"Yes",
					"No",
					"Internal",
					"External",
					"Media"
				]
			}
		}`), edit_awarenessToBeMade); err != nil {
			return err
		}
		collection.Schema.AddField(edit_awarenessToBeMade)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("67689ye74w7y9pa")
		if err != nil {
			return err
		}

		// update
		edit_awarenessToBeMade := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "xap6irdp",
			"name": "awarenessToBeMade",
			"type": "select",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"Yes",
					"No",
					"Internal",
					"External",
					"Media"
				]
			}
		}`), edit_awarenessToBeMade); err != nil {
			return err
		}
		collection.Schema.AddField(edit_awarenessToBeMade)

		return dao.SaveCollection(collection)
	})
}
