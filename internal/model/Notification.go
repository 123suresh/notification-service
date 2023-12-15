package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	ID      uuid.UUID `json:"id"`
	Details string    `json:"details"`
}

type NotifyRequest struct {
	Details string `json:"details"`
}

type NotifyResponse struct {
	ID      uuid.UUID `json:"id"`
	Details string    `json:"details"`
}

func NewNotify(req *NotifyRequest) *Notification {
	return &Notification{
		ID:      uuid.New(),
		Details: req.Details,
	}
}

func (noti *Notification) NotifyRes() *NotifyResponse {
	return &NotifyResponse{
		ID:      noti.ID,
		Details: noti.Details,
	}
}
