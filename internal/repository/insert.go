package repository

import (
	"api/pkg/models"
	"context"
	"time"
)

func (d *DatabaseRepo) InsertToAccounts(account models.AccountReq) (id int, err error) {
	err = d.db.QueryRow(context.Background(), "INSERT INTO accounts (firstname, lastname, email, password) VALUES ($1, $2, $3, $4) RETURNING id", account.FirstName, account.LastName, account.Email, account.Password).Scan(&id)
	return
}

func (d *DatabaseRepo) InsertToLocations(latitude, longitude float64) (id int, err error) {
	err = d.db.QueryRow(context.Background(), `INSERT INTO locations (latitude, longitude) VALUES ($1, $2) RETURNING id`, latitude, longitude).Scan(&id)
	return
}

func (d *DatabaseRepo) InsertToTypes(animaltype string) (id int, err error) {
	err = d.db.QueryRow(context.Background(), `INSERT INTO types (type) VALUES ($1) RETURNING id`, animaltype).Scan(&id)
	return
}

func (d *DatabaseRepo) InsertToAnimals(animal models.Animal) (id int, err error) {
	err = d.db.QueryRow(context.Background(), `INSERT INTO animals (weight, length, height, gender, "chipperId", "chippingLocationId") VALUES($1,$2,$3,$4,$5,$6) RETURNING id`, animal.Weight, animal.Length, animal.Height, animal.Gender, animal.ChipperId, animal.ChippingLocationId).Scan(&id)
	return
}

func (d *DatabaseRepo) InsertToAnimalTypes(animalId, typeId int) (err error) {
	_, err = d.db.Exec(context.Background(), `INSERT INTO animal_types ("animalId", "typeId") VALUES ($1,$2)`, animalId, typeId)
	return
}

func (d *DatabaseRepo) InsertToVisitedLocation(animalId, locationId int) (id int, date time.Time, err error) {
	err = d.db.QueryRow(context.Background(), `INSERT INTO locations_visited  ("locationPointId", "animalId") VALUES ($1, $2) RETURNING id, date`, locationId, animalId).Scan(&id, &date)
	return
}
