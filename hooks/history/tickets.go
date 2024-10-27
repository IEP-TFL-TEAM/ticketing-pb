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

		return GenerateTicketNumber(app, e, authRecord)
	})
}

func NewTicketHistory(app *pocketbase.PocketBase) {
	app.OnRecordAfterCreateRequest("tickets", "comments").Add(func(e *core.RecordCreateEvent) error {
		authRecord, _ := e.HttpContext.Get(apis.ContextAuthRecordKey).(*models.Record)
		if authRecord == nil {
			return nil
		}

		if e.Record.Collection().Name == "tickets" {
			err := CreateTicketCreationHistory(app, e, authRecord)
			if err != nil {
				log.Fatal("Error creating ticket create history")
			}
		} else if e.Record.Collection().Name == "comments" {
			err := CreateCommentHistory(app, e, authRecord)
			if err != nil {
				log.Fatal(err)
			}
		}

		return nil
	})

}

func GenerateTicketNumber(app *pocketbase.PocketBase, e *core.RecordCreateEvent, authRecord *models.Record) error {
	ticketCount, err := app.Dao().FindRecordById("ticketCount", "1")
	println(ticketCount)

	if err != nil {
		return err
	}

	now := time.Now()
	datePart := now.Format("020106")
	timePart := now.Format("150405")

	ticketNumber := fmt.Sprintf("%s%s TT", datePart, timePart)

	e.Record.Set("ticketNumber", ticketNumber)
	e.Record.Set("count", ticketCount.GetString("totalItems"))

	return nil
}
