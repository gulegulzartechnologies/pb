package main

import (
	"log"
	"os"
	"strings"

	_ "github.com/gulegulzartechnologies/pb/migrations"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {

	app := pocketbase.New()

	// app.OnAfterBootstrap().Add(func(e *core.BootstrapEvent) error {
	// 	return nil
	// })

	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Admin UI
		// (the isGoRun check is to enable it only during development)
		Automigrate: isGoRun,
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

}
