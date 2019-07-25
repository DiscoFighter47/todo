package data

import "github.com/DiscoFighter47/todo/backend/model"

// UserStore ...
type UserStore interface {
	AddUser(*model.User) error
	GetUser(string) (*model.User, error)
}
