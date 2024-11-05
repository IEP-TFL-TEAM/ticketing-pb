package history

import (
	"fmt"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func NewCommentHistory(app *pocketbase.PocketBase) {
	app.OnRecordAfterCreateRequest("comments").Add(func(e *core.RecordCreateEvent) error {
		authRecord, _ := e.HttpContext.Get(apis.ContextAuthRecordKey).(*models.Record)
		if authRecord == nil {
			return nil
		}

		collection, err := app.Dao().FindCollectionByNameOrId("history")
		if err != nil {
			return err
		}

		ticket, _ := app.Dao().FindRecordById("tickets", e.Record.GetString("ticketId"))
		authFullName := fmt.Sprintf("%s %s",
			authRecord.GetString("firstName"),
			authRecord.GetString("lastName"))

		record := models.NewRecord(collection)
		record.Set("action", fmt.Sprintf("%s added a new comment on this ticket", authFullName))
		record.Set("ticketId", ticket.GetId())

		if err := app.Dao().SaveRecord(record); err != nil {
			return err
		}

		return nil
	})
}
