package history

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func NewTicket(app *pocketbase.PocketBase) {
	app.OnRecordBeforeCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		authRecord, _ := e.HttpContext.Get(apis.ContextAuthRecordKey).(*models.Record)
		if authRecord == nil {
			return nil
		}

		if e.Record.Collection().Name != "tickets" {
			return nil
		}

		return IncrementTicketCount(app, e, authRecord)
	})
}

func NewTicketHistory(app *pocketbase.PocketBase) {
	app.OnRecordAfterCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		authRecord, _ := e.HttpContext.Get(apis.ContextAuthRecordKey).(*models.Record)
		if authRecord == nil {
			return nil
		}

		if e.Record.Collection().Name == "tickets" {
			err := CreateTicketCreationHistory(app, e, authRecord)
			if err != nil {
				log.Fatal("Error creating ticket create history")
			}
		}

		if e.Record.Collection().Name == "comments" {
			err := CreateCommentHistory(app, e, authRecord)
			if err != nil {
				log.Fatal(err)
			}
		}

		return nil
	})
}

func IncrementTicketCount(app *pocketbase.PocketBase, e *core.RecordCreateEvent, authRecord *models.Record) error {
	ticketCount, err := app.Dao().FindRecordById("ticketCount", "1")
	println(ticketCount)

	if err != nil {
		return err
	}
	e.Record.Set("count", ticketCount.GetString("totalItems"))

	return nil
}
