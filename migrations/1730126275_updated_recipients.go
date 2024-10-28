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

		collection, err := dao.FindCollectionByNameOrId("c607t2ux8dbejkb")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("941xvvg8")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("c607t2ux8dbejkb")
		if err != nil {
			return err
		}

		// add
		del_categoryId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "941xvvg8",
			"name": "categoryId",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "yrgvxehupjjixov",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), del_categoryId); err != nil {
			return err
		}
		collection.Schema.AddField(del_categoryId)

		return dao.SaveCollection(collection)
	})
}
