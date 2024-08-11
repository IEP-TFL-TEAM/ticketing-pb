package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("t0djc6rrp2vtymy")
		if err != nil {
			return err
		}

		collection.Name = "categorylevels"

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("t0djc6rrp2vtymy")
		if err != nil {
			return err
		}

		collection.Name = "severitylevels"

		return dao.SaveCollection(collection)
	})
}
