package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("yv3cl86s3yci1co")
		if err != nil {
			return err
		}

		collection.Name = "changeteams"

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("yv3cl86s3yci1co")
		if err != nil {
			return err
		}

		collection.Name = "changeteam"

		return dao.SaveCollection(collection)
	})
}
