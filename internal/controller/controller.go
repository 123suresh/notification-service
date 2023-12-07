package controller

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/suresh/notification/internal/service"
)

type Controller struct {
	router *gin.Engine
	svc    service.ServiceInterface
}

func NewController(svc service.ServiceInterface) *Controller {
	ctl := &Controller{}
	ctl.router = gin.Default()
	ctl.svc = svc
	return ctl
}

func (ctl *Controller) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}
	err := ctl.router.Run(":" + port)
	if err != nil {
		log.Fatalf("Unable to run server on port %s", port)
	}
}
