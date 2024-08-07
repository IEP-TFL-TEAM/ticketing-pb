package tickets

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func IncrementTicketCount(app *pocketbase.PocketBase, e *core.RecordCreateEvent) error {
	ticketCount, err := app.Dao().FindRecordById("ticketCount", "1")

	if err != nil {
		return err
	}
	e.Record.Set("count", ticketCount.GetString("totalItems"))

	return nil
}
