package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("rybge8g8z4jgrwi")
		if err != nil {
			return err
		}

		collection.DeleteRule = types.Pointer("")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("rybge8g8z4jgrwi")
		if err != nil {
			return err
		}

		collection.DeleteRule = nil

		return dao.SaveCollection(collection)
	})
}
