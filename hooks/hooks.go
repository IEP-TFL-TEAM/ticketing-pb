package hooks

import (
	"github.com/pocketbase/pocketbase"
	"ticketing-pb/hooks/notifications"
	"ticketing-pb/hooks/tickets"
)

func BindCustomHooks(app *pocketbase.PocketBase) {
	tickets.BindCustomHooks(app)
	notifications.BindCustomHooks(app)
}
