package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("4fgpahxwll5pzuj")
		if err != nil {
			return err
		}

		collection.Name = "teamequipments"

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("4fgpahxwll5pzuj")
		if err != nil {
			return err
		}

		collection.Name = "teamEquipments"

		return dao.SaveCollection(collection)
	})
}
