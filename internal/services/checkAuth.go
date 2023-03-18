package services

func (s *Services) CheckAuth(login, password string) bool {

	hash := s.passwordToHash(password)

	count, err := s.Database.db.SelectEmailAndPassword(login, hash)
	if err != nil || count == 0 {
		return false
	}

	return true
}
