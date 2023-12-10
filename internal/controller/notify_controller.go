package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/suresh/notification/internal/model"
	"github.com/suresh/notification/utils/response"
)

func (ctl *Controller) CreateNotify(c *gin.Context) {
	logrus.Info("data => ")
	notifyReq := &model.NotifyRequest{}
	logrus.Info("data => ", notifyReq)
	err := c.ShouldBindJSON(notifyReq)
	if err != nil {
		logrus.Error("json bind error :: ", err)
		response.ERROR(c, err, http.StatusBadRequest)
		return
	}
	notifyResponse, code, err := ctl.svc.CreateNotify(notifyReq)
	if err != nil {
		response.ERROR(c, err, code)
		return
	}
	response.JSON(c, notifyResponse, "Success", 0, 0)
}
