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

		collection, err := dao.FindCollectionByNameOrId("caza2oj0rtxw1gu")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("@request.auth.id != \"\"")

		collection.ViewRule = types.Pointer("@request.auth.id != \"\"")

		collection.CreateRule = types.Pointer("@request.auth.id != \"\"")

		collection.UpdateRule = types.Pointer("@request.auth.id != \"\"")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("caza2oj0rtxw1gu")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("@request.auth.collectionName = \"users\"")

		collection.ViewRule = types.Pointer("@request.auth.collectionName = \"users\"")

		collection.CreateRule = types.Pointer("@request.auth.id != \"\" && @request.auth.role != \"staff\"")

		collection.UpdateRule = types.Pointer("@request.auth.id != \"\" && @request.auth.role != \"staff\"")

		return dao.SaveCollection(collection)
	})
}
