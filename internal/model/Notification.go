package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	ID uuid.UUID `json:"id"`
}
