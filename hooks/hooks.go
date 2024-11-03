package hooks

import (
	"github.com/pocketbase/pocketbase"
	"ticketing-pb/hooks/history"
)

func BindCustomHooks(app *pocketbase.PocketBase) {
	history.BindCustomHistory(app)
}
