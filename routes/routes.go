package routes

import (
	"github.com/pocketbase/pocketbase"
)

func BindCustomRoutes(app *pocketbase.PocketBase) {
	SendBroadcastEmail(app)
	SendRoutineCreationNotification(app)
	SendRequestCreationNotification(app)
}
