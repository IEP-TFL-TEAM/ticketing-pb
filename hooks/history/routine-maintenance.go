package history

import (
	"fmt"
	"time"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func NewRoutineMaintenance(app *pocketbase.PocketBase) {
	app.OnRecordBeforeCreateRequest("routinemaintenance").Add(func(e *core.RecordCreateEvent) error {
		authRecord, _ := e.HttpContext.Get(apis.ContextAuthRecordKey).(*models.Record)
		if authRecord == nil {
			return nil
		}

		return GenerateRoutineMaintenanceNumber(app, e, authRecord)
	})
}

func GenerateRoutineMaintenanceNumber(app *pocketbase.PocketBase, e *core.RecordCreateEvent, authRecord *models.Record) error {
	now := time.Now()
	ticketNumber := fmt.Sprintf("%s%s RM", now.Format("020106"), now.Format("150405"))

	e.Record.Set("ticketNumber", ticketNumber)

	return nil
}
