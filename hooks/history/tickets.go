package history

import (
	"fmt"
	"reflect"
	"time"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func NewTicket(app *pocketbase.PocketBase) {
	app.OnRecordBeforeCreateRequest("tickets").Add(func(e *core.RecordCreateEvent) error {
		authRecord, _ := e.HttpContext.Get(apis.ContextAuthRecordKey).(*models.Record)
		if authRecord == nil {
			return nil
		}

		now := time.Now()
		ticketNumber := fmt.Sprintf("%s%s TT", now.Format("020106"), now.Format("150405"))

		e.Record.Set("ticketNumber", ticketNumber)

		return nil
	})
}

func NewTicketHistory(app *pocketbase.PocketBase) {
	app.OnRecordBeforeCreateRequest("tickets").Add(func(e *core.RecordCreateEvent) error {
		authRecord, _ := e.HttpContext.Get(apis.ContextAuthRecordKey).(*models.Record)
		if authRecord == nil {
			return nil
		}

		collection, err := app.Dao().FindCollectionByNameOrId("history")
		if err != nil {
			return err
		}

		authFullName := fmt.Sprintf("%s %s",
			authRecord.GetString("firstName"),
			authRecord.GetString("lastName"))

		record := models.NewRecord(collection)
		record.Set("action", fmt.Sprintf("%s created a new ticket", authFullName))
		record.Set("ticketId", e.Record.Id)

		if err := app.Dao().SaveRecord(record); err != nil {
			return err
		}

		return nil
	})
}

func TrackTicketFieldUpdates(app *pocketbase.PocketBase) {
	app.OnRecordAfterUpdateRequest("tickets").Add(func(e *core.RecordUpdateEvent) error {
		authRecord, _ := e.HttpContext.Get(apis.ContextAuthRecordKey).(*models.Record)
		if authRecord == nil {
			return nil
		}

		record := e.Record
		oldRecord := record.OriginalCopy()

		authFullName := fmt.Sprintf("%s %s",
			authRecord.GetString("firstName"),
			authRecord.GetString("lastName"))

		collection, err := app.Dao().FindCollectionByNameOrId("history")
		if err != nil {
			return fmt.Errorf("failed to find history collection: %v", err)
		}

		fieldsToCheck := []string{"status", "title", "description", "categoryId", "incidentStart"}

		for _, field := range fieldsToCheck {
			oldValue := oldRecord.Get(field)
			newValue := record.Get(field)

			if !reflect.DeepEqual(oldValue, newValue) {
				historyRecord := models.NewRecord(collection)
				actionDetail := fmt.Sprintf("%s updated the %s from %v to %v for this ticket", authFullName, field, oldValue, newValue)

				if field == "incidentStart" {
					actionDetail = fmt.Sprintf("%s updated the Incident Start Time for this ticket", authFullName)
				} else if field == "categoryId" {
					actionDetail = fmt.Sprintf("%s updated the Category and Severity for this ticket", authFullName)
				}

				historyRecord.Set("action", actionDetail)
				historyRecord.Set("ticketId", record.Id)

				if err := app.Dao().SaveRecord(historyRecord); err != nil {
					return fmt.Errorf("failed to save history record: %v", err)
				}

				fmt.Printf("Change recorded: Field '%s' in record %s\n", field, record.Id)
			}
		}

		return nil
	})
}
