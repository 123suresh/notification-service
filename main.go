package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/suresh/notification/internal/controller"
	"github.com/suresh/notification/internal/grpc"
	"github.com/suresh/notification/internal/repository"
	"github.com/suresh/notification/internal/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("Error getting env, not coming through %v", err)
	}
	logrus.Info("Successfully loaded env file")
	repo := repository.NewRepo()
	svc := service.NewService(repo)
	ctl := controller.NewController(svc)
	grpc.CreateServer()
	ctl.Run()
}
