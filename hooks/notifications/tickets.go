package notifications

import (
	"fmt"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func NewTicket(app *pocketbase.PocketBase, e *core.RecordCreateEvent) error {
	categoryId := fmt.Sprintf("%v", e.Record.Get("categoryId"))

	mailDetails := MailDetails{
		toMailAddresses: GetMailAddressesForAutoEmail(app, categoryId),
		subject:         "New Ticket",
		sendTo:          "Team",
		action:          "Created",
		url:             app.Settings().Meta.AppUrl + "/tickets/" + string(e.Record.Id),
	}

	return SendEmailNotification(app, mailDetails)
}

func ResolveTicket(app *pocketbase.PocketBase, e *core.RecordUpdateEvent) error {
	categoryId := fmt.Sprintf("%v", e.Record.Get("categoryId"))

	mailDetails := MailDetails{
		toMailAddresses: GetMailAddressesForAutoEmail(app, categoryId),
		subject:         "Ticket Pending",
		sendTo:          "Team",
		action:          "re-opened. Please resolve",
		url:             app.Settings().Meta.AppUrl + "/tickets/" + string(e.Record.Id),
	}

	return SendEmailNotification(app, mailDetails)
}

func ApprovedTicket(app *pocketbase.PocketBase, e *core.RecordUpdateEvent) error {
	categoryId := fmt.Sprintf("%v", e.Record.Get("categoryId"))

	mailDetails := MailDetails{
		toMailAddresses: GetMailAddressesForAutoEmail(app, categoryId),
		subject:         "Ticket Approved",
		sendTo:          "Team",
		action:          "approved",
		url:             app.Settings().Meta.AppUrl + "/tickets/" + string(e.Record.Id),
	}

	return SendEmailNotification(app, mailDetails)
}

func ClosedTicket(app *pocketbase.PocketBase, e *core.RecordUpdateEvent) error {
	categoryId := fmt.Sprintf("%v", e.Record.Get("categoryId"))

	mailDetails := MailDetails{
		toMailAddresses: GetMailAddressesForAutoEmail(app, categoryId),
		subject:         "Ticket Closed",
		sendTo:          "Team",
		action:          "closed",
		url:             app.Settings().Meta.AppUrl + "/tickets/" + string(e.Record.Id),
	}

	return SendEmailNotification(app, mailDetails)
}
