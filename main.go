package main

import (
	"log"
	"os"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	"ticketing-pb/hooks"
	_ "ticketing-pb/migrations"
	"ticketing-pb/routes"
)

func main() {
	app := pocketbase.New()

	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	// Enable auto creation of migration files when making collection changes in the Admin UI
	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: isGoRun,
	})

	hooks.BindCustomHooks(app)
	routes.BindCustomRoutes(app)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
