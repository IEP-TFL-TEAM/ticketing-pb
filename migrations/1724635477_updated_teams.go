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

		collection.CreateRule = types.Pointer("@request.auth.collectionName = \"users\" && @request.auth.role = \"admin\"")

		collection.UpdateRule = types.Pointer("@request.auth.collectionName = \"users\" && @request.auth.role = \"admin\"")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("rybge8g8z4jgrwi")
		if err != nil {
			return err
		}

		collection.CreateRule = types.Pointer("@request.auth.collectionName = \"users\"")

		collection.UpdateRule = types.Pointer("@request.auth.collectionName = \"users\"")

		return dao.SaveCollection(collection)
	})
}
