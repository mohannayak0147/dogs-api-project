package service

import (
	"dogs-common/config"
	"dogs-common/repository"
)

type Service struct {
	Conn repository.DogRepository
	Cfg  config.Provider
}

func NewAPIService(cfg config.Provider) (*Service, error) {
	db, err := repository.NewDB(cfg.GetString("DB_User"), cfg.GetString("DB_Password"), cfg.GetString("DB_Host"), cfg.GetString("DB_Database"), cfg.GetInt("DB_Port"))
	if err != nil {
		return nil, err
	}

	return &Service{
		Conn: db,
		Cfg:  cfg,
	}, nil
}
