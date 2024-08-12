package notifications

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

var err = godotenv.Load()
var uiURL = os.Getenv("UI_URL")

func CreateTicketEmailNotification(app *pocketbase.PocketBase, e *core.RecordCreateEvent, authRecord *models.Record) error {
	mailDetails := MailDetails{
		toMailAddresses: GetMailAddressesByRole(app, "admin"),
		subject:         "New Ticket",
		sendTo:          "Test Room",
		action:          "Created",
		url:             uiURL + "/tickets/" + string(e.Record.Id),
	}

	return SendEmailNotification(app, mailDetails)
}
