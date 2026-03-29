package dogs

import (
	"dogs-api/service"
	"dogs-api/v1/http"
	"dogs-common/model"
	"dogs-common/repository"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

var DogSVC *service.Service

// GetDog :
// @Summary      Returns a dog
// @Description  Fetch a dog by name.
// @Tags         Dogs
// @Param        dog path string true "Dog breed"
// @Produce      json
// @Success      200 {object} model.Dog "dog"
// @Failure      500 {object} routes.ErrorResp "Error description"
// @Failure      404 {object} routes.ErrorResp "could not retrieve dog : dog not found"
// @Router       /api/dogs/{dog} [get]
func GetDog(c *fiber.Ctx) error {
	log.Infof("Get dog request for %s", c.Params("dog"))
	requestedDog := c.Params("dog")
	var dog model.Dog
	var err error
	dog, err = DogSVC.Conn.RetrieveDog(requestedDog)
	if err != nil {
		log.Errorf("could not retrieve dog: %v", err)
		if errors.Is(err, repository.ErrDogNotFound) {
			return http.NewError(fmt.Errorf("could not retrieve dog: %w", err), fiber.StatusNotFound)
		}
		return http.NewError(fmt.Errorf("could not retrieve dog: %w", err), fiber.StatusInternalServerError)
	}

	log.Infof("Successfully completed dog request for %s", c.Params("dog"))

	return c.JSON(dog)
}

// GetDogs
// @Summary      Returns the list of all dogs
// @Description  Provides dogs information
// @Tags         Dogs
// @Produce      json
// @Success      200 {object} []model.Dog "dogs"
// @Failure      500 {object} routes.ErrorResp "Error description"
// @Router       /api/dogs [get]
func GetDogs(c *fiber.Ctx) error {
	log.Info("Get dogs request ")
	dogs, err := DogSVC.Conn.RetrieveAllDogs()
	if err != nil {
		log.Errorf("could not retrieve dogs : %v", err)
		return http.NewError(fmt.Errorf("could not retrieve dogs : %w", err), fiber.StatusInternalServerError)
	}

	log.Info("Successfully completed get dogs request")
	return c.JSON(dogs)
}

// CreateDog
// @Summary      Create a new dog
// @Description  Creates a new dog entry.
// @Tags         Dogs
// @Accept       json
// @Produce      json
// @Param        dog body model.Dog true "Dog payload"
// @Success      201 {object} model.Dog "created dog"
// @Failure      400 {object} routes.ErrorResp "invalid request body"
// @Failure      500 {object} routes.ErrorResp "could not create dog"
// @Router       /api/dogs [post]
func CreateDog(c *fiber.Ctx) error {
	log.Info("Create dog request")

	var dog model.Dog
	if err := c.BodyParser(&dog); err != nil {
		return http.NewError(fmt.Errorf("invalid request body: %w", err), fiber.StatusBadRequest)
	}

	if err := DogSVC.Conn.CreateNewDog(dog); err != nil {
		return http.NewError(fmt.Errorf("could not create dog: %w", err), fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusCreated).JSON(dog)
}

// UpdateDog
// @Summary      Update an existing dog
// @Description  Updates a dog by breed.
// @Tags         Dogs
// @Accept       json
// @Produce      json
// @Param        dog path string true "Dog breed"
// @Param        body body model.Dog true "Updated dog payload"
// @Success      200 {object} model.Dog "updated dog"
// @Failure      400 {object} routes.ErrorResp "Invalid request body"
// @Failure      500 {object} routes.ErrorResp "Internal server error"
// @Failure      404 {object} routes.ErrorResp "Dog not found"
// @Router       /api/dogs/{dog} [put]
func UpdateDog(c *fiber.Ctx) error {
	breed := c.Params("dog")

	var dog model.Dog
	if err := c.BodyParser(&dog); err != nil {
		return http.NewError(fmt.Errorf("invalid request body: %w", err), fiber.StatusBadRequest)
	}

	rows, err := DogSVC.Conn.UpdateDog(breed, dog)
	if err != nil {
		return http.NewError(fmt.Errorf("could not update dog: %w", err), fiber.StatusInternalServerError)
	}

	if rows == 0 {
		return http.NewError(repository.ErrDogNotFound, fiber.StatusNotFound)
	}

	updatedDog, err := DogSVC.Conn.RetrieveDog(dog.Breed)
	if err != nil {
		return http.NewError(fmt.Errorf("could not retrieve updated dog: %w", err), fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(updatedDog)
}

// DeleteDog
// @Summary      Delete a dog
// @Description  Deletes a dog by breed.
// @Tags         Dogs
// @Produce      json
// @Param        dog path string true "Dog breed"
// @Success      204 "No Content"
// @Failure      404 {object} routes.ErrorResp "Dog not found"
// @Failure      500 {object}  routes.ErrorResp "Internal server error"
// @Router       /api/dogs/{dog} [delete]
func DeleteDog(c *fiber.Ctx) error {
	breed := c.Params("dog")

	rows, err := DogSVC.Conn.RemoveExistingDog(breed)
	if err != nil {
		return http.NewError(fmt.Errorf("could not delete dog: %w", err), fiber.StatusInternalServerError)
	}

	if rows == 0 {
		return http.NewError(repository.ErrDogNotFound, fiber.StatusNotFound)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
