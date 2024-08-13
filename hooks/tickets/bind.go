package tickets

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func BindCustomHooks(app *pocketbase.PocketBase) {
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

	app.OnRecordAfterUpdateRequest().Add(func(e *core.RecordUpdateEvent) error {
		authRecord, _ := e.HttpContext.Get(apis.ContextAuthRecordKey).(*models.Record)
		if authRecord == nil {
			return nil
		}

		if e.Record.Collection().Name == "tickets" {
			err := CreateTicketUpdateHistory(app, e, authRecord)
			if err != nil {
				log.Fatal("Error creating comment create history")
			}
		}

		return nil
	})
}
