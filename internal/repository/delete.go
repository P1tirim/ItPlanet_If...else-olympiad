package repository

import "context"

func (d *DatabaseRepo) DeleteAccount(id int) (count int, err error) {
	com, err := d.db.Exec(context.Background(), `DELETE FROM accounts WHERE id = $1`, id)
	if err != nil {
		return
	}

	count = int(com.RowsAffected())
	return
}

func (d *DatabaseRepo) DeleteLocation(id int) (count int, err error) {
	com, err := d.db.Exec(context.Background(), `DELETE FROM locations WHERE id = $1`, id)
	if err != nil {
		return
	}

	count = int(com.RowsAffected())
	return
}

func (d *DatabaseRepo) DeleteType(id int) (count int, err error) {
	com, err := d.db.Exec(context.Background(), `DELETE FROM types WHERE id = $1`, id)
	if err != nil {
		return
	}

	count = int(com.RowsAffected())
	return
}

func (d *DatabaseRepo) DeleteVisitedLocation(visitedLocationId, animalId int) (count int, err error) {
	com, err := d.db.Exec(context.Background(), `DELETE FROM locations_visited WHERE id = $1 AND "animalId" = $2`, visitedLocationId, animalId)
	if err != nil {
		return
	}

	count = int(com.RowsAffected())
	return
}

func (d *DatabaseRepo) DeleteAnimal(id int) (count int, err error) {
	com, err := d.db.Exec(context.Background(), `DELETE FROM animals WHERE id = $1`, id)
	if err != nil {
		return
	}

	count = int(com.RowsAffected())
	return
}

func (d *DatabaseRepo) DeleteTypeOfAnimal(animalId, typeId int) (count int, err error) {
	com, err := d.db.Exec(context.Background(), `DELETE FROM animal_types WHERE "animalId" = $1 AND "typeId" = $2`, animalId, typeId)
	if err != nil {
		return
	}

	count = int(com.RowsAffected())
	return
}
