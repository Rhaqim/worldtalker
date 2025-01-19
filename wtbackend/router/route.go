package router

import (
	"github.com/Rhaqim/wtbackend/service"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
	wsService *service.WebsocketService
}

func NewRouter(wsService *service.WebsocketService) *Router {
	rGin := gin.Default()

	router := &Router{
		Engine:    rGin,
		wsService: wsService,
	}

	router.registerRoutes()

	return router
}

func (r *Router) registerRoutes() {
	r.GET("/ws/:id", r.wsService.Handle)
}

func (r *Router) Run(addr string) error {
	return r.Engine.Run(addr)
}
