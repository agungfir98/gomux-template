package router

import (
	"github.com/agungfir98/gomux-starter/api"
	"github.com/agungfir98/gomux-starter/handler"
)

func ExampleRoute(api *api.ApiHandler) {
	api.Get("/", handler.ExampleHandler)
}
