package notifications

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

var err = godotenv.Load()
var uiURL = os.Getenv("UI_URL")

func NewTicket(app *pocketbase.PocketBase, e *core.RecordCreateEvent) error {
	mailDetails := MailDetails{
		toMailAddresses: GetMailAddressesForAutoEmail(app),
		subject:         "New Ticket",
		sendTo:          "Team",
		action:          "Created",
		url:             uiURL + "/tickets/" + string(e.Record.Id),
	}

	return SendEmailNotification(app, mailDetails)
}

func ResolveTicket(app *pocketbase.PocketBase, e *core.RecordUpdateEvent) error {
	mailDetails := MailDetails{
		toMailAddresses: GetMailAddressesForAutoEmail(app),
		subject:         "Ticket Pending",
		sendTo:          "Team",
		action:          "re-opened. Please resolve",
		url:             uiURL + "/tickets/" + string(e.Record.Id),
	}

	return SendEmailNotification(app, mailDetails)
}

func ApprovedTicket(app *pocketbase.PocketBase, e *core.RecordUpdateEvent) error {
	mailDetails := MailDetails{
		toMailAddresses: GetMailAddressesForAutoEmail(app),
		subject:         "Ticket Approved",
		sendTo:          "Team",
		action:          "approved",
		url:             uiURL + "/tickets/" + string(e.Record.Id),
	}

	return SendEmailNotification(app, mailDetails)
}

func ClosedTicket(app *pocketbase.PocketBase, e *core.RecordUpdateEvent) error {
	mailDetails := MailDetails{
		toMailAddresses: GetMailAddressesForAutoEmail(app),
		subject:         "Ticket Closed",
		sendTo:          "Team",
		action:          "closed",
		url:             uiURL + "/tickets/" + string(e.Record.Id),
	}

	return SendEmailNotification(app, mailDetails)
}
