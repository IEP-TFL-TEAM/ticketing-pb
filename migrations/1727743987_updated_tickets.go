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

		// update
		edit_incidentStart := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "edqlsddc",
			"name": "incidentStart",
			"type": "date",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), edit_incidentStart); err != nil {
			return err
		}
		collection.Schema.AddField(edit_incidentStart)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("caza2oj0rtxw1gu")
		if err != nil {
			return err
		}

		// update
		edit_incidentStart := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "edqlsddc",
			"name": "startDate",
			"type": "date",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), edit_incidentStart); err != nil {
			return err
		}
		collection.Schema.AddField(edit_incidentStart)

		return dao.SaveCollection(collection)
	})
}
