package repository

import (
	"database/sql"
	"dogs-common/model"
	"dogs-common/storage"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -destination=./rate_mock.go -package=repository dogs-common/repository DogRepository
type DogRepository interface {
	CreateNewDog(dog model.Dog) error
	UpdateDog(dogBreed string, dog model.Dog) (int64, error)
	RemoveExistingDog(dogBreed string) (int64, error)
	RetrieveAllDogs() ([]model.Dog, error)
	RetrieveDog(dogBreed string) (model.Dog, error)
	Close()
}

type DBConn struct {
	DB *sqlx.DB
}

var ErrDogNotFound = fmt.Errorf("dog not found")

func NewDB(user, password, host, dbname string, port int) (*DBConn, error) {
	db, err := storage.GetDB(user, password, host, dbname, port)
	if err != nil {
		return nil, fmt.Errorf("could not instantiate dog repository %w", err)
	}
	return &DBConn{DB: db}, nil
}

// CreateNewDog : this creates a new record in dog table for a particular dog
func (d *DBConn) CreateNewDog(dog model.Dog) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	jsonString, err := json.Marshal(dog.Data)
	if err != nil {
		return fmt.Errorf("cant marshal %v", err)
	}

	queryStr := `INSERT INTO dogs (breed,data,updated_at,created_at) VALUES (?,?,?,?)`
	_, err = d.DB.Exec(queryStr,
		dog.Breed,
		string(jsonString),
		now,
		now)
	if err != nil {
		return fmt.Errorf("could not insert new record for dog [%v] : %w", dog, err)
	}

	return nil
}

// UpdateDog : it updates dog details
func (d *DBConn) UpdateDog(dogBreed string, dog model.Dog) (int64, error) {
	now := time.Now().Format("2006-01-02T15:04:05")
	jsonString, err := json.Marshal(dog.Data)
	if err != nil {
		return 0, fmt.Errorf("cant marshal %v", err)
	}

	query := "UPDATE dogs SET breed=?, data=?, updated_at=? WHERE breed=?"

	res, err := d.DB.Exec(query, dog.Breed, string(jsonString), now, dogBreed)
	if err != nil {
		return 0, fmt.Errorf("unable to execute query : %w", err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("could not get numbers of rows affected : %s", err.Error())
	}

	return rows, nil
}

// RemoveExistingDog : this deletes existing record in dog table for a particular breed
func (d *DBConn) RemoveExistingDog(dogBreed string) (int64, error) {
	res, err := d.DB.Exec("DELETE FROM dogs WHERE breed = ?", dogBreed)
	if err != nil {
		return 0, fmt.Errorf("could not remove existing dog[%s] : %w", dogBreed, err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("could not get rows affected: %w", err)
	}
	return rows, nil
}

// RetrieveDog : it retrieves a dog
func (d *DBConn) RetrieveDog(dogBreed string) (model.Dog, error) {
	var dog model.Dog
	var jsonStr string

	err := d.DB.QueryRow(`SELECT breed, data FROM dogs where breed = ?`, dogBreed).Scan(&dog.Breed, &jsonStr)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Dog{}, ErrDogNotFound
		}
		return model.Dog{}, fmt.Errorf("query failed: %w", err)
	}

	err = json.Unmarshal([]byte(jsonStr), &dog.Data)
	if err != nil {
		return model.Dog{}, fmt.Errorf("unable to unmarshall [%s] : %w", jsonStr, err)
	}

	return dog, nil
}

// RetrieveAllDogs : it retrieves all dogs
func (d *DBConn) RetrieveAllDogs() ([]model.Dog, error) {
	var results []model.Dog

	rows, err := d.DB.Query(`SELECT breed, data FROM dogs`)
	if err != nil {
		return results, fmt.Errorf("unable to execute query : %w", err)
	}
	for rows.Next() {
		var dog model.Dog
		var jsonStr string

		err := rows.Scan(&dog.Breed, &jsonStr)
		if err != nil {
			return results, fmt.Errorf("scan failed : %w", err)
		}

		err = json.Unmarshal([]byte(jsonStr), &dog.Data)
		if err != nil {
			return results, fmt.Errorf("unable to unmarshall [%s] : %w", jsonStr, err)
		}

		results = append(results, dog)
	}
	if rows.Err() != nil {
		return results, fmt.Errorf("row iteration failure: %w", err)
	}
	rows.Close()
	return results, nil
}

// Close : to close the database connection
func (d *DBConn) Close() {
	d.DB.Close()
}
