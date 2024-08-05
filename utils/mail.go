package utils

import (
	"net/mail"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
)

func GetMailAddressesByCollection(app *pocketbase.PocketBase, collection string) []mail.Address {
	baseUsers, _ := app.Dao().FindRecordsByExpr(collection,
		dbx.HashExp{"verified": true},
	)

	toMailAddresses := []mail.Address{}
	for _, baseUser := range baseUsers {
		mailAddress := mail.Address{
			Address: baseUser.Email(),
			Name:    baseUser.GetString("firstName") + " " + baseUser.GetString("lastName"),
		}
		toMailAddresses = append(toMailAddresses, mailAddress)
	}

	return toMailAddresses
}
