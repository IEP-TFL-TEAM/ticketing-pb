package tickets

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func IncrementTicketCount(app *pocketbase.PocketBase, e *core.RecordCreateEvent, authRecord *models.Record) error {
	ticketCount, err := app.Dao().FindRecordById("ticketCount", "1")
	println(ticketCount)

	if err != nil {
		return err
	}
	e.Record.Set("count", ticketCount.GetString("totalItems"))

	return nil
}
