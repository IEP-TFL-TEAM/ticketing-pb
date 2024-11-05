package history

import (
	"fmt"
	"log"
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
	app.OnRecordAfterCreateRequest("tickets").Add(func(e *core.RecordCreateEvent) error {
		authRecord, _ := e.HttpContext.Get(apis.ContextAuthRecordKey).(*models.Record)
		if authRecord == nil {
			return nil
		}

		err := CreateTicketCreationHistory(app, e, authRecord)
		if err != nil {
			log.Fatal("Error creating ticket create history")
		}

		return nil
	})
}
