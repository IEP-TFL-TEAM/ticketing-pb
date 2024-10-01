package history

import (
	"github.com/pocketbase/pocketbase"
)

func BindCustomHooks(app *pocketbase.PocketBase) {
	NewTicket(app)
	NewTicketHistory(app)
	NewCommentHistory(app)

	NewChangeRequest(app)
	NewRoutineMaintenance(app)
}
