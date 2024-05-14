package database

import "net/url"

func NewConnString(driver, username, password, host, db string) string {
	u := url.URL{
		Scheme: driver,
		User:   url.UserPassword(username, password),
		Host:   host,
		Path:   db,
	}

	return u.String()
}
