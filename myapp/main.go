package main

import (
	"myapp/handers"

	"github.com/tsawler/celeritas"
)

type application struct {
	App     *celeritas.Celeritas
	Handers *handers.Handers
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
