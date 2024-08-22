package notifications

import (
	"fmt"
	"log"
	"net/mail"
	"os"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/tools/mailer"
)

func GetMailAddressesByTeamId(app *pocketbase.PocketBase, teamId string) []mail.Address {
	baseUsers, _ := app.Dao().FindRecordsByExpr("users",
		dbx.HashExp{"teamId": teamId},
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

func GetMailAddressesForAutoEmail(app *pocketbase.PocketBase) []mail.Address {
	baseUsers, _ := app.Dao().FindRecordsByExpr("recipients",
		dbx.HashExp{"verified": true},
	)

	toMailAddresses := []mail.Address{}
	for _, baseUser := range baseUsers {
		mailAddress := mail.Address{
			Address: baseUser.Email(),
			Name:    baseUser.GetString("name"),
		}
		toMailAddresses = append(toMailAddresses, mailAddress)
	}

	return toMailAddresses
}

func ParseAddressFromEmail(email string) mail.Address {
	return mail.Address{Address: email}
}

func SendEmailNotification(app *pocketbase.PocketBase, mailDetails MailDetails) error {
	fileContent, err := os.ReadFile("templates/ticket.html")
	if err != nil {
		log.Fatal(err)
	}

	nocMail := os.Getenv("NOC")

	mailBody := string(fileContent)

	cc := []mail.Address{
		ParseAddressFromEmail(nocMail),
	}

	message := &mailer.Message{
		From: mail.Address{
			Address: app.Settings().Meta.SenderAddress,
			Name:    app.Settings().Meta.SenderName,
		},
		To:      mailDetails.toMailAddresses,
		Subject: mailDetails.subject,
		HTML: fmt.Sprintf(mailBody,
			mailDetails.sendTo,
			mailDetails.action,
			mailDetails.url,
		),
		Cc: cc,
	}
	return app.NewMailClient().Send(message)
}

type MailDetails struct {
	toMailAddresses []mail.Address
	subject         string
	sendTo          string
	action          string
	url             string
}
