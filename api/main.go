package main

import (
	_ "dogs-api/docs"
	"dogs-api/service"
	routes "dogs-api/v1"
	"dogs-api/v1/dogs"
	"dogs-api/v1/health"
	"dogs-common/config"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title         Dog API
// @description   Dog API
// @version 1.0
// @host          127.0.0.1:8081
// @BasePath      /
func main() {
	log.Info("Starting the application")
	app := fiber.New(fiber.Config{ErrorHandler: routes.StandardErrorHandler})
	cfg := config.Config()
	svc, err := service.NewAPIService(cfg)
	if err != nil {
		log.Fatalf("could not create service %v", err)
	}
	defer svc.Conn.Close()

	dogs.DogSVC = svc

	v1 := app.Group("/api")
	// Read
	v1.Get("/status", health.HealthCheckHandler)

	v1.Get("/dogs", dogs.GetDogs)
	v1.Get("/dogs/:dog", dogs.GetDog)

	app.Get("/docs/*", fiberSwagger.WrapHandler)

	// CREATE
	v1.Post("/dogs", dogs.CreateDog)

	// UPDATE
	v1.Put("/dogs/:dog", dogs.UpdateDog)

	// DELETE
	v1.Delete("/dogs/:dog", dogs.DeleteDog)

	err = app.Listen(cfg.GetString("API_PORT"))
	if err != nil {
		log.Fatalf(err.Error())
	}
}
