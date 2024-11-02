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

		// remove
		collection.Schema.RemoveField("g3ia97dz")

		// add
		new_status := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "wqk3nun8",
			"name": "status",
			"type": "select",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"PENDING",
					"CLOSED"
				]
			}
		}`), new_status); err != nil {
			return err
		}
		collection.Schema.AddField(new_status)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("67689ye74w7y9pa")
		if err != nil {
			return err
		}

		// add
		del_isClosed := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "g3ia97dz",
			"name": "isClosed",
			"type": "bool",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), del_isClosed); err != nil {
			return err
		}
		collection.Schema.AddField(del_isClosed)

		// remove
		collection.Schema.RemoveField("wqk3nun8")

		return dao.SaveCollection(collection)
	})
}
