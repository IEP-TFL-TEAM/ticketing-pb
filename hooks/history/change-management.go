package history

import (
	"fmt"
	"time"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func NewChangeRequest(app *pocketbase.PocketBase) {
	app.OnRecordBeforeCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		authRecord, _ := e.HttpContext.Get(apis.ContextAuthRecordKey).(*models.Record)
		if authRecord == nil {
			return nil
		}

		if e.Record.Collection().Name != "changerequests" {
			return nil
		}

		return GenerateChangeRequestNumber(app, e, authRecord)
	})
}

func GenerateChangeRequestNumber(app *pocketbase.PocketBase, e *core.RecordCreateEvent, authRecord *models.Record) error {
	now := time.Now()
	datePart := now.Format("020106")
	timePart := now.Format("150405")

	ticketNumber := fmt.Sprintf("%s%s CR", datePart, timePart)

	e.Record.Set("ticketNumber", ticketNumber)

	return nil
}
