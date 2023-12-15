package service

import (
	"net/http"

	"github.com/suresh/notification/internal/model"
)

type Notify interface {
	CreateNotify(notifyReq *model.NotifyRequest) (notifyRes *model.NotifyResponse, code int, err error)
}

func (svc *Service) CreateNotify(notifyReq *model.NotifyRequest) (notifyRes *model.NotifyResponse, code int, err error) {
	noti := model.NewNotify(notifyReq)
	result, err := svc.repo.CreateNotify(noti)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	response := result.NotifyRes()
	return response, http.StatusCreated, nil
}
