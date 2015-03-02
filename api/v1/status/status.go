package status

import (
	"github.com/cornfeedhobo/gin-boiler-plate/api/v1"

	"github.com/gin-gonic/gin"
)

func RealMain(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
}

func init() {
	v1.RegisterGroup(RealMain)
}
