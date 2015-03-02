package v1

import (
	"github.com/cornfeedhobo/gin-boiler-plate/api"

	"github.com/gin-gonic/gin"
)

var groupRegFuncs []func(*gin.RouterGroup)

func RegisterGroup(regFunc func(*gin.RouterGroup)) {
	groupRegFuncs = append(groupRegFuncs, regFunc)
}

func RealMain(engine *gin.Engine) {
	group := engine.Group("/v1")
	for _, regFunc := range groupRegFuncs {
		regFunc(group)
	}
}

func init() {
	api.RegisterHandler(RealMain)
}
