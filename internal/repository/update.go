package repository

import (
	"api/pkg/models"
	"context"
)

func (d *DatabaseRepo) UpdateAccount(id int, acc models.AccountReq) error {
	_, err := d.db.Exec(context.Background(), "UPDATE accounts SET firstname = $1, lastname = $2, email = $3, password = $4 WHERE id = $5", acc.FirstName, acc.LastName, acc.Email, acc.Password, id)
	return err
}

func (d *DatabaseRepo) UpdateLocations(id int, latitude, longitude float64) (count int, err error) {
	com, err := d.db.Exec(context.Background(), `UPDATE locations SET latitude = $1, longitude = $2 WHERE id = $3`, latitude, longitude, id)
	if err != nil {
		return
	}

	count = int(com.RowsAffected())
	return
}

func (d *DatabaseRepo) UpdateType(id int, animalType string) (count int, err error) {
	com, err := d.db.Exec(context.Background(), `UPDATE types SET type = $1 WHERE id = $2`, animalType, id)
	if err != nil {
		return
	}

	count = int(com.RowsAffected())
	return
}

func (d *DatabaseRepo) UpdateVisitedLocation(visitedLocationId, animalId, locationId int) (count int, err error) {
	com, err := d.db.Exec(context.Background(), `UPDATE locations_visited SET "locationPointId" = $3 WHERE id = $1 AND "animalId" = $2`, visitedLocationId, animalId, locationId)
	if err != nil {
		return
	}

	count = int(com.RowsAffected())
	return
}

func (d *DatabaseRepo) UpdateAniaml(animal models.Animal) (err error) {
	_, err = d.db.Exec(context.Background(), `UPDATE animals SET weight = $1, length = $2, height = $3, gender = $4, "lifeStatus" = $5, "chipperId" = $6, "chippingLocationId" = $7 WHERE id = $8`, animal.Weight, animal.Length, animal.Height, animal.Gender, animal.LifeStatus, animal.ChipperId, animal.ChippingLocationId, animal.Id)
	if err != nil {
		return
	}
	return
}

func (d *DatabaseRepo) UpdateDeathTimeAnimal(animalId int) (err error) {
	_, err = d.db.Exec(context.Background(), `UPDATE animals SET "deathDateTime" = now() WHERE id = $1`, animalId)
	return
}

func (d *DatabaseRepo) UpdateTypeOfAnimal(animalId, oldType, newType int) (err error) {
	_, err = d.db.Exec(context.Background(), `UPDATE animal_types SET "typeId" = $1 WHERE "animalId" = $2 AND "typeId" = $3`, newType, animalId, oldType)
	return
}
