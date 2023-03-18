package services

import "api/pkg/models"

func (s *Services) GetAccountById(id int) (models.AccountResp, error) {
	return s.Database.db.SelectAccount(id)
}

func (s *Services) GetAccountSearch(firstName, lastName, email string, from, size int) ([]models.AccountResp, error) {
	return s.Database.db.SelectAccoutnSearch(firstName, lastName, email, from, size)
}

func (s *Services) GetIdAccountByEmail(email string) (int, error) {
	return s.Database.db.SelectIdAccountByEmail(email)
}

func (s *Services) GetAnimalById(id int) (models.Animal, error) {
	return s.Database.db.SelectAnimalById(id)
}

func (s *Services) GetAnimalSearch(search models.AnimalSearch) ([]models.Animal, error) {
	return s.Database.db.SelectAnimalSearch(search)
}

func (s *Services) GetTypeOfAnimal(idType int) (int, string, error) {
	return s.Database.db.SelectTypeOfAnimal(idType)
}

func (s *Services) GetLocations(idLocations int) (int, float64, float64, error) {
	return s.Database.db.SelectLocation(idLocations)
}

func (s *Services) GetVisitedLocations(idAnimal int, startDate string, endDate string, from int, size int) ([]models.VisitedLocations, error) {
	return s.Database.db.SelectVisitedLocations(idAnimal, startDate, endDate, from, size)
}

func (s *Services) GetTypesOfAnimal(idAnimal int) ([]int, error) {
	return s.Database.db.SelectTypesOfAnimal(idAnimal)
}

func (s *Services) IsAccountInAnimal(idAccount int) (bool, error) {
	count, err := s.Database.db.SelectAccountInAnimals(idAccount)
	if err != nil || count != 0 {
		return true, err
	}

	return false, nil
}

func (s *Services) IsDistintcCoordinates(latitude, longitude float64) (bool, error) {
	count, err := s.Database.db.SelectLocationByCoordinates(latitude, longitude)
	if err != nil || count != 0 {
		return false, err
	}

	return true, nil
}

// Return false if location isn't in animals
func (s *Services) IsLocationInAnimals(idLocation int) (bool, error) {
	count, err := s.Database.db.SelectLocationInAnimals(idLocation)
	if err != nil || count != 0 {
		return true, err
	}

	return false, nil
}

func (s *Services) IsDistinctType(animalType string) (bool, error) {
	count, err := s.Database.db.SelectTypesByType(animalType)
	if err != nil || count != 0 {
		return false, err
	}

	return true, nil
}

func (s *Services) IsTypeInAnimals(idType int) (bool, error) {
	count, err := s.Database.db.SelectTypesInAnimal(idType)
	if err != nil || count != 0 {
		return true, err
	}

	return false, nil
}

func (s *Services) IsLocationInVisitedLocation(locationId int) (bool, error) {
	count, err := s.Database.db.SelectCountVisitedLocation(locationId)
	if err != nil || count != 0 {
		return true, err
	}

	return false, nil
}

func (s *Services) GetVisitedLocationsAll(animalId int) ([]models.VisitedLocations, error) {
	return s.Database.db.SelectVisitedLocationsAll(animalId)
}
