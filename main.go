package main

import (
	"github.com/cornfeedhobo/gin-boiler-plate/api"

	_ "github.com/cornfeedhobo/gin-boiler-plate/api/v1/status"
)

func main() {
	app := api.New()
	app.Router.Run("127.0.0.1:8080")
}
