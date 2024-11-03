package history

import (
	"fmt"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
)

func CreateHistoryRecord(app *pocketbase.PocketBase, collectionName string, ticketId string, action string) error {
	collection, err := app.Dao().FindCollectionByNameOrId(collectionName)
	if err != nil {
		return err
	}

	record := models.NewRecord(collection)
	form := forms.NewRecordUpsert(app, record)

	form.LoadData(map[string]any{
		"ticketId": ticketId,
		"action":   action,
	})

	return form.Submit()
}

func CreateCommentHistory(app *pocketbase.PocketBase, e *core.RecordCreateEvent, authRecord *models.Record) error {
	collectionName := "history"

	ticket, _ := app.Dao().FindRecordById("tickets", e.Record.GetString("ticketId"))

	action := fmt.Sprintf("%s commented on ticket #%s", authRecord.GetString("firstName"), ticket.GetString("count"))

	return CreateHistoryRecord(app, collectionName, ticket.GetId(), action)
}

func CreateTicketCreationHistory(app *pocketbase.PocketBase, e *core.RecordCreateEvent, authRecord *models.Record) error {
	collectionName := "history"

	ticketId := e.Record.Id

	action := fmt.Sprintf("%s created ticket #%s", authRecord.GetString("firstName"), e.Record.GetString("count"))

	return CreateHistoryRecord(app, collectionName, ticketId, action)
}

func CreateTicketUpdateHistory(app *pocketbase.PocketBase, e *core.RecordUpdateEvent, authRecord *models.Record) error {
	collectionName := "history"

	ticketId := e.Record.Id

	fieldsToCheck := []string{"status", "title", "description", "reportedBy", "categoryId"}

	for _, field := range fieldsToCheck {
		if e.Record.Get(field) != e.Record.OriginalCopy().Get(field) {

			originalStatus, newStatus := e.Record.OriginalCopy().GetString(field), e.Record.GetString(field)

			if field == "categoryId" {
				field = "Category"

				if originalStatus != "" {
					oldRecord, _ := app.Dao().FindRecordById("categories", originalStatus)
					originalStatus = oldRecord.GetString("name")
				}

				newRecord, _ := app.Dao().FindRecordById("categories", newStatus)
				newStatus = newRecord.GetString("name")
			}

			actionMsg := fmt.Sprintf("%s changed %s", authRecord.GetString("firstName"), field)
			if originalStatus != "" {
				actionMsg += fmt.Sprintf(" from %s to %s", originalStatus, newStatus)
			} else {
				actionMsg += fmt.Sprintf(" to %s", newStatus)
			}

			actionMsg += fmt.Sprintf(" for ticket #%s", e.Record.GetString("count"))

			err := CreateHistoryRecord(app, collectionName, ticketId, actionMsg)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
