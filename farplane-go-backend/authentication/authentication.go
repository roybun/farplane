package authentication

import "errors"

type userCredentials struct {
	username string
	password string
}

var hardcodedCredentials = [2]userCredentials{
	{"User1", "PW1"},
	{"User2", "Pw2"},
}

func userExists(username string) (int, error) {
	for index, userCredentials := range hardcodedCredentials {
		if userCredentials.username == username {
			return index, nil
		}
	}
	return -1, errors.New("User does not exist")
}

type FPAuthenticator interface {
	ValidateUser(username string, password string) (string, error)
}

type HardcodedAuthenticator struct {
}

func (HardcodedAuthenticator) ValidateUser(username string, password string) error {
	index, ok := userExists(username)
	if ok == nil {
		if hardcodedCredentials[index].password == password {
			return nil
		}
	}
	return errors.New("CAnnot validate user")
}

type PostgresAuthenticator struct {
}

func (PostgresAuthenticator) ValidateUser(username string, password string) error {
	return errors.New("Postgres Authentication not implemented yet!")
}
