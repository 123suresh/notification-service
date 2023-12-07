package service

import "github.com/suresh/notification/internal/repository"

type Service struct {
	repo repository.RepoInterface
}

type ServiceInterface interface {
}

func NewService(repo repository.RepoInterface) ServiceInterface {
	svc := &Service{}
	svc.repo = repo
	svc.repo = repository.NewRepo()
	return svc
}
