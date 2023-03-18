package repository

import (
	"api/pkg/models"
	"context"
	"fmt"
	"log"
	"time"
)

func (d *DatabaseRepo) CountEmailInDB(id int, email string) (count int, err error) {
	err = d.db.QueryRow(context.Background(), "SELECT COUNT(*) FROM accounts WHERE email = $1 AND id <> $2", email, id).Scan(&count)
	return
}

func (d *DatabaseRepo) SelectAccount(id int) (acc models.AccountResp, err error) {
	rows, err := d.db.Query(context.Background(), "SELECT id, firstname, lastname, email FROM accounts WHERE id = $1", id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&acc.ID, &acc.FirstName, &acc.LastName, &acc.Email)
		return
	}

	return
}

func (d *DatabaseRepo) SelectEmailAndPassword(email, password string) (count int, err error) {

	err = d.db.QueryRow(context.Background(), "SELECT COUNT(*) FROM accounts WHERE email = $1 AND password = $2", email, password).Scan(&count)

	return
}

func (d *DatabaseRepo) SelectAccoutnSearch(firstName, lastName, email string, from, size int) ([]models.AccountResp, error) {
	firstName = "%" + firstName + "%"
	lastName = "%" + lastName + "%"
	email = "%" + email + "%"

	rows, err := d.db.Query(context.Background(), "SELECT id, firstname, lastname, email FROM accounts WHERE firstname ILIKE $1 AND lastname ILIKE $2 AND email ILIKE $3 ORDER BY id ASC LIMIT $4 OFFSET $5", firstName, lastName, email, size, from)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.AccountResp

	for rows.Next() {
		var acc models.AccountResp

		err = rows.Scan(&acc.ID, &acc.FirstName, &acc.LastName, &acc.Email)
		if err != nil {
			continue
		}

		result = append(result, acc)
	}

	return result, nil
}

func (d *DatabaseRepo) SelectIdAccountByEmail(email string) (id int, err error) {
	rows, err := d.db.Query(context.Background(), "SELECT id FROM accounts WHERE email = $1", email)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id)
		return
	}
	return
}

func (d *DatabaseRepo) SelectAnimalById(id int) (animal models.Animal, err error) {
	rows, err := d.db.Query(context.Background(), "SELECT * FROM animals WHERE id = $1", id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&animal.Id, &animal.Weight, &animal.Length, &animal.Height, &animal.Gender, &animal.LifeStatus, &animal.ChippingDateTime, &animal.ChipperId, &animal.ChippingLocationId, &animal.DeathDateTime)
		if err != nil {
			return
		}
	}

	animal.AnimalTypes, err = d.SelectTypesOfAnimal(animal.Id)
	if err != nil {
		return
	}

	if animal.AnimalTypes == nil {
		animal.AnimalTypes = []int{}
	}

	animal.VisitedLocations, err = d.selectVisitedLocationsId(animal.Id)

	if animal.VisitedLocations == nil {
		animal.VisitedLocations = []int{}
	}

	return
}

func (d *DatabaseRepo) SelectAnimalSearch(search models.AnimalSearch) (animals []models.Animal, err error) {
	where := ""
	if search.ChipperId != 0 {
		where += fmt.Sprintf(` "chipperId" = %d `, search.ChipperId)
	}

	if search.ChippingLocationId != 0 {
		if where != "" {
			where += "AND"
		}
		where += fmt.Sprintf(` "chippingLocationId" = %d `, search.ChippingLocationId)
	}

	if search.LifeStatus != "" {
		if where != "" {
			where += "AND"
		}
		where += fmt.Sprintf(` "lifeStatus" = '%s' `, search.LifeStatus)
	}

	if search.Gender != "" {
		if where != "" {
			where += "AND"
		}
		where += fmt.Sprintf(` "gender" = '%s' `, search.Gender)
	}

	if search.StartDateTime != "" && search.EndDateTime != "" {
		if where != "" {
			where += "AND"
		}
		where += fmt.Sprintf(` "chippingDateTime" BETWEEN '%s' AND '%s'`, search.StartDateTime, search.EndDateTime)
	} else if search.StartDateTime != "" {
		if where != "" {
			where += "AND"
		}
		where += fmt.Sprintf(` "chippingDateTime" > '%s'`, search.StartDateTime)
	} else if search.EndDateTime != "" {
		if where != "" {
			where += "AND"
		}
		where += fmt.Sprintf(` "chippingDateTime" < '%s'`, search.EndDateTime)
	}

	if where != "" {
		where = "WHERE" + where
	}

	rows, err := d.db.Query(context.Background(), `SELECT * FROM animals `+where+` ORDER BY id LIMIT $1 OFFSET $2`, search.Size, search.From)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var animal models.Animal

		err = rows.Scan(&animal.Id, &animal.Weight, &animal.Length, &animal.Height, &animal.Gender, &animal.LifeStatus, &animal.ChippingDateTime, &animal.ChipperId, &animal.ChippingLocationId, &animal.DeathDateTime)
		if err != nil {
			log.Println(err)
			continue
		}

		animal.AnimalTypes, err = d.SelectTypesOfAnimal(animal.Id)
		if err != nil {
			log.Println(err)
			continue
		}

		if animal.AnimalTypes == nil {
			animal.AnimalTypes = []int{}
		}

		animal.VisitedLocations, err = d.selectVisitedLocationsId(animal.Id)
		if err != nil {
			log.Println(err)
			continue
		}

		if animal.VisitedLocations == nil {
			animal.VisitedLocations = []int{}
		}

		animals = append(animals, animal)
	}

	return
}

func (d *DatabaseRepo) SelectTypesOfAnimal(idAnimal int) (types []int, err error) {
	rows, err := d.db.Query(context.Background(), `SELECT "typeId" FROM animal_types WHERE "animalId" = $1`, idAnimal)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var typeId int

		err = rows.Scan(&typeId)
		if err != nil {
			log.Println(err)
			continue
		}

		types = append(types, typeId)
	}

	return
}

func (d *DatabaseRepo) selectVisitedLocationsId(idAnimal int) (locations []int, err error) {
	rows, err := d.db.Query(context.Background(), `SELECT id FROM locations_visited WHERE "animalId" = $1`, idAnimal)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var locationId int

		err = rows.Scan(&locationId)
		if err != nil {
			return
		}

		locations = append(locations, locationId)
	}

	return
}

func (d *DatabaseRepo) SelectTypeOfAnimal(idType int) (id int, animalType string, err error) {
	rows, err := d.db.Query(context.Background(), "SELECT * FROM types WHERE id = $1", idType)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &animalType)
		return
	}

	return
}

func (d *DatabaseRepo) SelectLocation(idLocation int) (id int, latitude, longtitude float64, err error) {
	rows, err := d.db.Query(context.Background(), "SELECT * FROM locations WHERE id = $1", idLocation)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &latitude, &longtitude)
		return
	}

	return
}

func (d *DatabaseRepo) SelectVisitedLocations(idAnimal int, startDate, endDate string, from, size int) (result []models.VisitedLocations, err error) {
	where := fmt.Sprintf(`WHERE "animalId" = %d`, idAnimal)
	if startDate != "" && endDate != "" {
		where += fmt.Sprintf(" AND date BETWEEN '%s' AND '%s'", startDate, endDate)
	} else if startDate != "" {
		where += fmt.Sprintf(" AND date > '%s'", startDate)
	} else if endDate != "" {
		where += fmt.Sprintf(" AND date < '%s'", endDate)
	}

	rows, err := d.db.Query(context.Background(), `SELECT id, date, "locationPointId" FROM locations_visited `+where+` LIMIT $1 OFFSET $2`, size, from)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var location models.VisitedLocations
		err = rows.Scan(&location.Id, &location.Date, &location.LocationId)
		if err != nil {
			log.Println(err)
			continue
		}

		ch := location.Date.Format(time.RFC3339)
		location.Date, _ = time.Parse(time.RFC3339, ch)

		result = append(result, location)
	}

	return
}

func (d *DatabaseRepo) SelectVisitedLocationsAll(idAnimal int) (result []models.VisitedLocations, err error) {
	rows, err := d.db.Query(context.Background(), `SELECT id, date,"locationPointId" FROM locations_visited WHERE "animalId" = $1 ORDER BY id`, idAnimal)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var location models.VisitedLocations
		err = rows.Scan(&location.Id, &location.Date, &location.LocationId)
		if err != nil {
			log.Println(err)
			continue
		}

		result = append(result, location)
	}

	return
}

func (d *DatabaseRepo) SelectAccountInAnimals(idAccount int) (count int, err error) {
	err = d.db.QueryRow(context.Background(), `SELECT COUNT(*) FROM animals WHERE "chipperId" = $1`, idAccount).Scan(&count)
	return
}

func (d *DatabaseRepo) SelectLocationByCoordinates(latitude, longitude float64) (count int, err error) {
	err = d.db.QueryRow(context.Background(), `SELECT COUNT(*) FROM locations WHERE latitude = $1 AND longitude = $2`, latitude, longitude).Scan(&count)
	return
}

func (d *DatabaseRepo) SelectLocationInAnimals(idLocation int) (count int, err error) {
	err = d.db.QueryRow(context.Background(), `SELECT COUNT(*) FROM animals WHERE "chippingLocationId" = $1`, idLocation).Scan(&count)
	return
}

func (d *DatabaseRepo) SelectTypesByType(animalType string) (count int, err error) {
	err = d.db.QueryRow(context.Background(), `SELECT COUNT(*) FROM types WHERE type = $1`, animalType).Scan(&count)
	return
}

func (d *DatabaseRepo) SelectTypesInAnimal(idType int) (count int, err error) {
	err = d.db.QueryRow(context.Background(), `SELECT COUNT(*) FROM animal_types WHERE "typeId" = $1`, idType).Scan(&count)
	return
}

func (d *DatabaseRepo) SelectCountVisitedLocation(locationId int) (count int, err error) {
	err = d.db.QueryRow(context.Background(), `SELECT COUNT(*) FROM locations_visited WHERE "locationPointId" = $1`, locationId).Scan(&count)
	return
}
