package services

import (
	"api/pkg/models"
	"time"
)

func (s *Services) AddLocation(latitude, longititude float64) (int, error) {
	return s.Database.db.InsertToLocations(latitude, longititude)
}

func (s *Services) AddType(animalType string) (int, error) {
	return s.Database.db.InsertToTypes(animalType)
}

func (s *Services) AddAnimal(animal models.Animal) (models.Animal, error) {
	animalId, err := s.Database.db.InsertToAnimals(animal)
	if err != nil {
		return models.Animal{}, err
	}

	for _, animaltype := range animal.AnimalTypes {
		err = s.Database.db.InsertToAnimalTypes(animalId, animaltype)
	}

	return s.Database.db.SelectAnimalById(animalId)
}

func (s *Services) AddVisitedLocation(animalId, locationId int) (int, time.Time, error) {
	return s.Database.db.InsertToVisitedLocation(animalId, locationId)
}

func (s *Services) AddTypeToAnimals(animalId, animalType int) (err error) {
	return s.Database.db.InsertToAnimalTypes(animalId, animalType)
}
