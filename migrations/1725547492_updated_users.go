package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		collection.CreateRule = types.Pointer("@request.auth.id != \"\"")

		collection.UpdateRule = types.Pointer("@request.auth.id != \"\"")

		options := map[string]any{}
		if err := json.Unmarshal([]byte(`{
			"allowEmailAuth": true,
			"allowOAuth2Auth": true,
			"allowUsernameAuth": true,
			"exceptEmailDomains": null,
			"manageRule": "@request.auth.id != \"\" && @request.auth.role = \"admin\"",
			"minPasswordLength": 8,
			"onlyEmailDomains": null,
			"onlyVerified": false,
			"requireEmail": false
		}`), &options); err != nil {
			return err
		}
		collection.SetOptions(options)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		collection.CreateRule = types.Pointer("@request.auth.role = \"admin\"")

		collection.UpdateRule = types.Pointer("@request.auth.role = \"admin\"")

		options := map[string]any{}
		if err := json.Unmarshal([]byte(`{
			"allowEmailAuth": true,
			"allowOAuth2Auth": true,
			"allowUsernameAuth": true,
			"exceptEmailDomains": null,
			"manageRule": "id = @request.auth.id && @request.auth.role = \"admin\"",
			"minPasswordLength": 8,
			"onlyEmailDomains": null,
			"onlyVerified": false,
			"requireEmail": false
		}`), &options); err != nil {
			return err
		}
		collection.SetOptions(options)

		return dao.SaveCollection(collection)
	})
}
