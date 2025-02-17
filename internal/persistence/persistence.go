package persistence

import (
	"errors"
	"net/url"
)

var (
	ErrUsernameAlreadyExists = errors.New("username already exists")
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrNotFound              = errors.New("user not found")
)

func NewConnString(driver, username, password, host, db string) string {
	u := url.URL{
		Scheme: driver,
		User:   url.UserPassword(username, password),
		Host:   host,
		Path:   db,
	}

	return u.String()
}
