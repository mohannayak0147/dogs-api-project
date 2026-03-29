package service

import (
	"dogs-common/config"
	"dogs-common/model"
	"dogs-common/repository"
	"encoding/json"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

type Service struct {
	conn repository.DogRepository
}

func NewService() (*Service, error) {
	cfg := config.Config()

	db, err := repository.NewDB(cfg.GetString("DB_User"), cfg.GetString("DB_Password"), cfg.GetString("DB_Host"), cfg.GetString("DB_Database"), cfg.GetInt("DB_Port"))
	if err != nil {
		return nil, err
	}

	return &Service{
		conn: db,
	}, nil
}

type DogData map[string][]string

func readJSON() ([]model.Dog, error) {
	file, err := os.Open("dogs.json")
	if err != nil {
		return nil, fmt.Errorf("could not open file[dogs.json]: %w", err)
	}
	defer file.Close()

	var dogData DogData
	if err := json.NewDecoder(file).Decode(&dogData); err != nil {
		return nil, fmt.Errorf("could not decode json: %w", err)
	}

	var dogs []model.Dog
	for breed, data := range dogData {
		dogs = append(dogs, model.Dog{
			Breed: breed,
			Data:  data,
		})
	}

	return dogs, nil
}

func (s *Service) Process() error {
	dogs, err := readJSON()
	if err != nil {
		return fmt.Errorf("could not parse input json file: %w", err)
	}

	for _, dog := range dogs {
		err = s.conn.CreateNewDog(dog)
		if err != nil {
			return err
		}
	}
	s.conn.Close()
	log.Infof("created %d new dogs", len(dogs))
	return nil
}
