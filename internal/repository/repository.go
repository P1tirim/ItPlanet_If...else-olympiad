package repository

import (
	"api/pkg/models"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database interface {
	CountEmailInDB(id int, email string) (count int, err error)
	SelectAccount(id int) (acc models.AccountResp, err error)
	SelectEmailAndPassword(email, password string) (count int, err error)
	SelectAccoutnSearch(firstName, lastName, email string, from, size int) ([]models.AccountResp, error)
	SelectIdAccountByEmail(email string) (id int, err error)
	SelectAnimalById(id int) (animal models.Animal, err error)
	SelectAnimalSearch(search models.AnimalSearch) (animals []models.Animal, err error)
	SelectTypesOfAnimal(idAnimal int) (types []int, err error)
	selectVisitedLocationsId(idAnimal int) (locations []int, err error)
	SelectTypeOfAnimal(idType int) (id int, animalType string, err error)
	SelectLocation(idLocation int) (id int, latitude, longtitude float64, err error)
	SelectVisitedLocations(idAnimal int, startDate, endDate string, from, size int) (result []models.VisitedLocations, err error)
	SelectVisitedLocationsAll(idAnimal int) (result []models.VisitedLocations, err error)
	SelectAccountInAnimals(idAccount int) (count int, err error)
	SelectLocationByCoordinates(latitude, longitude float64) (count int, err error)
	SelectLocationInAnimals(idLocation int) (count int, err error)
	SelectTypesByType(animalType string) (count int, err error)
	SelectTypesInAnimal(idType int) (count int, err error)
	SelectCountVisitedLocation(locationId int) (count int, err error)

	InsertToAccounts(account models.AccountReq) (id int, err error)
	InsertToLocations(latitude, longitude float64) (id int, err error)
	InsertToTypes(animaltype string) (id int, err error)
	InsertToAnimals(animal models.Animal) (id int, err error)
	InsertToAnimalTypes(animalId, typeId int) (err error)
	InsertToVisitedLocation(animalId, locationId int) (id int, date time.Time, err error)

	UpdateAccount(id int, acc models.AccountReq) error
	UpdateLocations(id int, latitude, longitude float64) (count int, err error)
	UpdateType(id int, animalType string) (count int, err error)
	UpdateVisitedLocation(visitedLocationId, animalId, locationId int) (count int, err error)
	UpdateAniaml(animal models.Animal) (err error)
	UpdateDeathTimeAnimal(animalId int) (err error)
	UpdateTypeOfAnimal(animalId, oldType, newType int) (err error)

	DeleteAccount(id int) (count int, err error)
	DeleteLocation(id int) (count int, err error)
	DeleteType(id int) (count int, err error)
	DeleteVisitedLocation(visitedLocationId, animalId int) (count int, err error)
	DeleteAnimal(id int) (count int, err error)
	DeleteTypeOfAnimal(animalId, typeId int) (count int, err error)
}

type Repository struct {
	Database Database
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{Database: newDatabaseRepo(db)}
}
