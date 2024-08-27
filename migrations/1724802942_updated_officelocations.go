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

		collection, err := dao.FindCollectionByNameOrId("7bfnibeo8zyr2si")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("q1sadm1y")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("7bfnibeo8zyr2si")
		if err != nil {
			return err
		}

		// add
		del_areaId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "q1sadm1y",
			"name": "areaId",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "63z0ygwtpjs28a4",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), del_areaId); err != nil {
			return err
		}
		collection.Schema.AddField(del_areaId)

		return dao.SaveCollection(collection)
	})
}
