package user

import (
	"context"
)

type service interface {
	FindByEmail(context.Context) error
}
