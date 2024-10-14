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

func OpenedTicket(app *pocketbase.PocketBase, e *core.RecordUpdateEvent) error {
	categoryId := fmt.Sprintf("%v", e.Record.Get("categoryId"))

	mailDetails := MailDetails{
		toMailAddresses: GetMailAddressesForAutoEmail(app, categoryId),
		subject:         "Ticket Opened",
		sendTo:          "Team",
		action:          "opened",
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

func EscalatedTicket(app *pocketbase.PocketBase, e *core.RecordUpdateEvent) error {
	categoryId := fmt.Sprintf("%v", e.Record.Get("categoryId"))

	mailDetails := MailDetails{
		toMailAddresses: GetMailAddressesForAutoEmail(app, categoryId),
		subject:         "Escalated Ticket",
		sendTo:          "Team",
		action:          "Escalated",
		url:             app.Settings().Meta.AppUrl + "/tickets/" + string(e.Record.Id),
	}

	return SendEmailNotification(app, mailDetails)
}
