package api

import "github.com/gin-gonic/gin"

var handlerRegFuncs []func(*gin.Engine)

func RegisterHandler(regFunc func(*gin.Engine)) {
	handlerRegFuncs = append(handlerRegFuncs, regFunc)
}

type Api struct {
	Router *gin.Engine
}

func New() *Api {
	api := &Api{
		Router: gin.Default(),
	}
	for _, regFunc := range handlerRegFuncs {
		regFunc(api.Router)
	}
	return api
}
