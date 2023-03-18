package services

import (
	"api/pkg/models"
)

func (s *Services) Registration(account models.AccountReq) (models.AccountResp, error) {

	account.Password = s.passwordToHash(account.Password)

	id, err := s.Database.db.InsertToAccounts(account)
	if err != nil {
		return models.AccountResp{}, err
	}

	return models.AccountResp{ID: id, FirstName: account.FirstName, LastName: account.LastName, Email: account.Email}, nil
}

func (s *Services) IsDistinctEmail(id int, email string) (bool, error) {
	count, err := s.Database.db.CountEmailInDB(id, email)
	if err != nil {
		return false, err
	}

	if count == 0 {
		return true, nil
	}

	return false, nil
}
