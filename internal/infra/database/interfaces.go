package database

import "github.com/pcechetto/go-api/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	Find(email string) (*entity.User, error)
} 