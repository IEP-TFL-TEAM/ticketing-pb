package notifications

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

var err = godotenv.Load()
var uiURL = os.Getenv("UI_URL")

func NewTicket(app *pocketbase.PocketBase) {
	app.OnRecordAfterCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		authRecord, _ := e.HttpContext.Get(apis.ContextAuthRecordKey).(*models.Record)
		if authRecord == nil {
			return nil
		}

		if e.Record.Collection().Name != "tickets" {
			return nil
		}

		return CreateTicketEmailNotification(app, e, authRecord)
	})
}

func CreateTicketEmailNotification(app *pocketbase.PocketBase, e *core.RecordCreateEvent, authRecord *models.Record) error {
	mailDetails := MailDetails{
		toMailAddresses: GetMailAddressesForAutoEmail(app),
		subject:         "New Ticket",
		sendTo:          "Team",
		action:          "Created",
		url:             uiURL + "/tickets/" + string(e.Record.Id),
	}

	return SendEmailNotification(app, mailDetails)
}
