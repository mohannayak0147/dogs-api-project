package main

import (
	"dogs-generator/service"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("started the application")
	svc, err := service.NewService()
	if err != nil {
		log.Fatalf("could not create service %v", err)
	}

	err = svc.Process()
	if err != nil {
		log.Fatalf("could not process dogs: %v", err)
	}
	log.Info("successfully finished processing dogs")
}
