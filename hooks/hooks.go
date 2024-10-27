package hooks

import (
	"github.com/pocketbase/pocketbase"
	"ticketing-pb/hooks/history"
	// "ticketing-pb/hooks/notifications"
)

func BindCustomHooks(app *pocketbase.PocketBase) {
	history.BindCustomHistory(app)
	// notifications.BindCustomNotifications(app)
}
