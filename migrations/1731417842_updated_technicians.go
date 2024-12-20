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

		collection, err := dao.FindCollectionByNameOrId("a3vplmoundlez9b")
		if err != nil {
			return err
		}

		// add
		new_email := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "hu6h5dpz",
			"name": "email",
			"type": "email",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"exceptDomains": [],
				"onlyDomains": []
			}
		}`), new_email); err != nil {
			return err
		}
		collection.Schema.AddField(new_email)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("a3vplmoundlez9b")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("hu6h5dpz")

		return dao.SaveCollection(collection)
	})
}
