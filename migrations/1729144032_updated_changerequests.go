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

		// add
		new_crSubmissionWithinFiveDays := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "5poybqyk",
			"name": "crSubmissionWithinFiveDays",
			"type": "select",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"Yes",
					"No"
				]
			}
		}`), new_crSubmissionWithinFiveDays); err != nil {
			return err
		}
		collection.Schema.AddField(new_crSubmissionWithinFiveDays)

		// add
		new_crDurationAdhered := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "1e3fnioc",
			"name": "crDurationAdhered",
			"type": "select",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"Yes",
					"No"
				]
			}
		}`), new_crDurationAdhered); err != nil {
			return err
		}
		collection.Schema.AddField(new_crDurationAdhered)

		// add
		new_crServiceImpactCorrect := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "2nfd5s2e",
			"name": "crServiceImpactCorrect",
			"type": "select",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"Yes",
					"No"
				]
			}
		}`), new_crServiceImpactCorrect); err != nil {
			return err
		}
		collection.Schema.AddField(new_crServiceImpactCorrect)

		// add
		new_crCorrectCustomerList := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "rhtukh9o",
			"name": "crCorrectCustomerList",
			"type": "select",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"Yes",
					"No"
				]
			}
		}`), new_crCorrectCustomerList); err != nil {
			return err
		}
		collection.Schema.AddField(new_crCorrectCustomerList)

		// add
		new_crTaskCompletion := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "paz8ujeu",
			"name": "crTaskCompletion",
			"type": "select",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"Yes",
					"No",
					"Partially Completed"
				]
			}
		}`), new_crTaskCompletion); err != nil {
			return err
		}
		collection.Schema.AddField(new_crTaskCompletion)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("67689ye74w7y9pa")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("5poybqyk")

		// remove
		collection.Schema.RemoveField("1e3fnioc")

		// remove
		collection.Schema.RemoveField("2nfd5s2e")

		// remove
		collection.Schema.RemoveField("rhtukh9o")

		// remove
		collection.Schema.RemoveField("paz8ujeu")

		return dao.SaveCollection(collection)
	})
}
