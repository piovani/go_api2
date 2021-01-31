package main

import (
	"github.com/piovani/go_api2/http"
	"github.com/piovani/go_api2/model"
)

var Products = model.Products{}

func main() {
	webserver := http.NewWebServer()

	webserver.Products = &Products

	webserver.Serve()
}
