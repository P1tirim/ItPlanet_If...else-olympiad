package services

import "api/pkg/models"

func (s *Services) UpdateAccountData(id int, acc models.AccountReq) error {
	acc.Password = s.passwordToHash(acc.Password)
	return s.Database.db.UpdateAccount(id, acc)
}

//Return true if found locations with this id
func (s *Services) UpdateLocations(id int, latitude, longitude float64) (bool, error) {
	count, err := s.Database.db.UpdateLocations(id, latitude, longitude)
	if err != nil || count == 0 {
		return false, err
	}

	return true, nil
}

func (s *Services) UpdateType(id int, animalType string) (bool, error) {
	count, err := s.Database.db.UpdateType(id, animalType)
	if err != nil || count == 0 {
		return false, err
	}

	return true, nil
}

func (s *Services) UpdateVisitedLocation(visitedLocationId, animalId, locationId int) (bool, error) {
	count, err := s.Database.db.UpdateVisitedLocation(visitedLocationId, animalId, locationId)
	if err != nil || count == 0 {
		return false, err
	}

	return true, nil
}

func (s *Services) UpdateAnimal(animal models.Animal) error {
	return s.Database.db.UpdateAniaml(animal)
}

func (s *Services) UpdateDeathTimeAnimal(animalId int) error {
	return s.Database.db.UpdateDeathTimeAnimal(animalId)
}

func (s *Services) UpdateTypeOfAnimal(animalId, oldType, newType int) error {
	return s.Database.db.UpdateTypeOfAnimal(animalId, oldType, newType)
}
