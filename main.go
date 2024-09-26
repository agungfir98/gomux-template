package main

import (
	"github.com/agungfir98/gomux-starter/api"
	"github.com/agungfir98/gomux-starter/router"
)

func main() {
	app := api.NewApiService(":8080")
	app.Use(router.ExampleRoute)

	app.Run()
}
