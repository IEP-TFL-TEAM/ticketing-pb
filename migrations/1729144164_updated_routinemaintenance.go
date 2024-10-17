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

		collection, err := dao.FindCollectionByNameOrId("w0xy9jxjhwjgnz3")
		if err != nil {
			return err
		}

		// add
		new_taskCompletion := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "2qtcz2em",
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
		}`), new_taskCompletion); err != nil {
			return err
		}
		collection.Schema.AddField(new_taskCompletion)

		// add
		new_alarmsCleared := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "q9oklhkq",
			"name": "alarmsCleared",
			"type": "select",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"Yes",
					"No",
					"Alarms Outstanding"
				]
			}
		}`), new_alarmsCleared); err != nil {
			return err
		}
		collection.Schema.AddField(new_alarmsCleared)

		// add
		new_serviceImpactCorrect := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "5l2w8ift",
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
		}`), new_serviceImpactCorrect); err != nil {
			return err
		}
		collection.Schema.AddField(new_serviceImpactCorrect)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("w0xy9jxjhwjgnz3")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("2qtcz2em")

		// remove
		collection.Schema.RemoveField("q9oklhkq")

		// remove
		collection.Schema.RemoveField("5l2w8ift")

		return dao.SaveCollection(collection)
	})
}
