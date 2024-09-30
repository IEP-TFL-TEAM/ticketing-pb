package notifications

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func BindCustomNotifications(app *pocketbase.PocketBase) {
	app.OnRecordAfterCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		authRecord, _ := e.HttpContext.Get(apis.ContextAuthRecordKey).(*models.Record)
		if authRecord == nil {
			return nil
		}

		if e.Record.Collection().Name != "tickets" {
			return nil
		}

		return NewTicket(app, e)
	})

	app.OnRecordAfterUpdateRequest().Add(func(e *core.RecordUpdateEvent) error {
		authRecord, _ := e.HttpContext.Get(apis.ContextAuthRecordKey).(*models.Record)
		if authRecord == nil {
			return nil
		}

		if e.Record.Collection().Name != "tickets" {
			return nil
		}

		fieldsToCheck := []string{"status"}

		for _, field := range fieldsToCheck {
			if e.Record.Get(field) != e.Record.OriginalCopy().Get(field) {

				newStatus := e.Record.GetString(field)

				if field == "status" {
					if newStatus == "PENDING" {
						return ResolveTicket(app, e)
					}
					if newStatus == "CLOSED" {
						return ClosedTicket(app, e)
					}
					if newStatus == "APPROVED" {
						return ApprovedTicket(app, e)
					}
				}
			}
		}
		return nil
	})
}
