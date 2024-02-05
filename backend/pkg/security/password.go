package security

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Password string

func (p Password) Hash() (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hashedPassword), nil
}

func (p Password) Matches(hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(p))
}
