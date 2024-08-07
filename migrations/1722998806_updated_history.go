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

		collection, err := dao.FindCollectionByNameOrId("jdq19kz788zjjg0")
		if err != nil {
			return err
		}

		// update
		edit_userId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "gdc74kyo",
			"name": "userId",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "_pb_users_auth_",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), edit_userId); err != nil {
			return err
		}
		collection.Schema.AddField(edit_userId)

		// update
		edit_ticketId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "xqxeca03",
			"name": "ticketId",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "caza2oj0rtxw1gu",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), edit_ticketId); err != nil {
			return err
		}
		collection.Schema.AddField(edit_ticketId)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("jdq19kz788zjjg0")
		if err != nil {
			return err
		}

		// update
		edit_userId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "gdc74kyo",
			"name": "user",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "_pb_users_auth_",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), edit_userId); err != nil {
			return err
		}
		collection.Schema.AddField(edit_userId)

		// update
		edit_ticketId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "xqxeca03",
			"name": "ticket",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "caza2oj0rtxw1gu",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), edit_ticketId); err != nil {
			return err
		}
		collection.Schema.AddField(edit_ticketId)

		return dao.SaveCollection(collection)
	})
}
