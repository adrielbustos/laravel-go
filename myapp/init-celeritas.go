package main

import (
	"log"
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
	// fmt.Println("DEBUG is set to", cel.Debug)
	cel.InfoLog.Println("DEBUG is set to", cel.Debug)

	app := &application{
		App: cel,
	}
	return app
}
