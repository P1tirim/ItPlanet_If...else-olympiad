package services

import "api/pkg/models"

func (s *Services) DeleteAccount(id int) (bool, error) {
	countAffected, err := s.Database.db.DeleteAccount(id)
	if err != nil || countAffected == 0 {
		return false, err
	}

	return true, nil
}

func (s *Services) DeleteLocaions(id int) (bool, error) {
	countAffected, err := s.Database.db.DeleteLocation(id)
	if err != nil || countAffected == 0 {
		return false, err
	}

	return true, nil
}

func (s *Services) DeleteType(id int) (bool, error) {
	countAffected, err := s.Database.db.DeleteType(id)
	if err != nil || countAffected == 0 {
		return false, err
	}

	return true, nil
}

func (s *Services) DeleteVisitedLocation(visitedLocationId int, animal models.Animal) (bool, error) {
	countAffected, err := s.Database.db.DeleteVisitedLocation(visitedLocationId, animal.Id)
	if err != nil || countAffected == 0 {
		return false, err
	}

	visitedLocations, err := s.GetVisitedLocationsAll(animal.Id)
	if err != nil {
		return false, err
	}

	if len(visitedLocations) > 0 && visitedLocations[0].LocationId == animal.ChippingLocationId {
		_, err = s.Database.db.DeleteVisitedLocation(visitedLocations[0].Id, animal.Id)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func (s *Services) DeleteAnimal(id int) (bool, error) {
	countAffected, err := s.Database.db.DeleteAnimal(id)
	if err != nil || countAffected == 0 {
		return false, err
	}

	return true, nil
}

func (s *Services) DeleteTypeOfAnimal(animalId, typeId int) (bool, error) {
	countAffected, err := s.Database.db.DeleteTypeOfAnimal(animalId, typeId)
	if err != nil || countAffected == 0 {
		return false, err
	}

	return true, nil
}
