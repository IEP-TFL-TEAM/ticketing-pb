package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("vvnj2k58yney9nn")
		if err != nil {
			return err
		}

		collection.Name = "region"

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("vvnj2k58yney9nn")
		if err != nil {
			return err
		}

		collection.Name = "regions"

		return dao.SaveCollection(collection)
	})
}
