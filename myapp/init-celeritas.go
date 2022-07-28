package main

import (
	"log"
	"myapp/handers"
	"os"

	"github.com/tsawler/celeritas"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init celeritas
	cel := &celeritas.Celeritas{}
	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "myapp"

	myHanders := &handers.Handers{
		App: cel,
	}

	app := &application{
		App:     cel,
		Handers: myHanders,
	}

	app.App.Routes = app.routes()

	return app
}
