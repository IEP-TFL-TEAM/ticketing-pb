package history

import (
	"github.com/pocketbase/pocketbase"
)

func BindCustomHistory(app *pocketbase.PocketBase) {
	NewTicket(app)
	NewTicketHistory(app)
	TrackTicketFieldUpdates(app)

	NewCommentHistory(app)

	NewChangeRequest(app)
	NewRoutineMaintenance(app)
}
