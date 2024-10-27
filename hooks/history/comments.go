package history

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func NewCommentHistory(app *pocketbase.PocketBase) {
	app.OnRecordAfterUpdateRequest("tickets").Add(func(e *core.RecordUpdateEvent) error {
		authRecord, _ := e.HttpContext.Get(apis.ContextAuthRecordKey).(*models.Record)
		if authRecord == nil {
			return nil
		}

		err := CreateTicketUpdateHistory(app, e, authRecord)
		if err != nil {
			log.Fatal("Error creating comment create history")
		}

		return nil
	})
}
