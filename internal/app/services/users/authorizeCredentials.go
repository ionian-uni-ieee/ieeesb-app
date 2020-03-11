package users

// AuthorizeCredentials authorizes username and password
func (s *Service) AuthorizeCredentials(username string, password string) (string, error) {
	usernameFilter := map[string]interface{}{
		"Username": username,
	}
	user, err := s.repositories.Users.FindOne(
		usernameFilter,
	)
	if err != nil {
		return "", err
	}
	isAuthorized, err := user.AuthorizePassword(password)
	if isAuthorized {
		return user.GetID(), err
	}

	return "", err
}
