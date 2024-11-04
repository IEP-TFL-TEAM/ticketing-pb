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

		collection, err := dao.FindCollectionByNameOrId("4fgpahxwll5pzuj")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("nshpmqil")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("4fgpahxwll5pzuj")
		if err != nil {
			return err
		}

		// add
		del_departmentId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "nshpmqil",
			"name": "departmentId",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "jwnpgoxuz954pjq",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), del_departmentId); err != nil {
			return err
		}
		collection.Schema.AddField(del_departmentId)

		return dao.SaveCollection(collection)
	})
}
