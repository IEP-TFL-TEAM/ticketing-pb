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
		edit_submissionWithinFiveDays := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "5poybqyk",
			"name": "submissionWithinFiveDays",
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
		}`), edit_submissionWithinFiveDays); err != nil {
			return err
		}
		collection.Schema.AddField(edit_submissionWithinFiveDays)

		// update
		edit_durationAdhered := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "1e3fnioc",
			"name": "durationAdhered",
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
		}`), edit_durationAdhered); err != nil {
			return err
		}
		collection.Schema.AddField(edit_durationAdhered)

		// update
		edit_serviceImpactCorrect := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "2nfd5s2e",
			"name": "serviceImpactCorrect",
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
		}`), edit_serviceImpactCorrect); err != nil {
			return err
		}
		collection.Schema.AddField(edit_serviceImpactCorrect)

		// update
		edit_correctCustomerList := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "rhtukh9o",
			"name": "correctCustomerList",
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
		}`), edit_correctCustomerList); err != nil {
			return err
		}
		collection.Schema.AddField(edit_correctCustomerList)

		// update
		edit_taskCompletion := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "paz8ujeu",
			"name": "taskCompletion",
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
		}`), edit_taskCompletion); err != nil {
			return err
		}
		collection.Schema.AddField(edit_taskCompletion)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("67689ye74w7y9pa")
		if err != nil {
			return err
		}

		// update
		edit_submissionWithinFiveDays := &schema.SchemaField{}
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
		}`), edit_submissionWithinFiveDays); err != nil {
			return err
		}
		collection.Schema.AddField(edit_submissionWithinFiveDays)

		// update
		edit_durationAdhered := &schema.SchemaField{}
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
		}`), edit_durationAdhered); err != nil {
			return err
		}
		collection.Schema.AddField(edit_durationAdhered)

		// update
		edit_serviceImpactCorrect := &schema.SchemaField{}
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
		}`), edit_serviceImpactCorrect); err != nil {
			return err
		}
		collection.Schema.AddField(edit_serviceImpactCorrect)

		// update
		edit_correctCustomerList := &schema.SchemaField{}
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
		}`), edit_correctCustomerList); err != nil {
			return err
		}
		collection.Schema.AddField(edit_correctCustomerList)

		// update
		edit_taskCompletion := &schema.SchemaField{}
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
		}`), edit_taskCompletion); err != nil {
			return err
		}
		collection.Schema.AddField(edit_taskCompletion)

		return dao.SaveCollection(collection)
	})
}
