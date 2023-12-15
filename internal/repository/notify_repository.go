package repository

import "github.com/suresh/notification/internal/model"

type NotifyInterface interface {
	CreateNotify(data *model.Notification) (*model.Notification, error)
}

func (r *Repo) CreateNotify(data *model.Notification) (*model.Notification, error) {
	err := r.db.Model(&model.Notification{}).Create(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
