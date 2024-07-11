package authcore

import (
	"github.com/tanasinp/go-inventory-management/database"
)

type UserRepository interface {
	CreateUser(user *database.User) error
	LoginUser(user *database.User) (string, error)
}
