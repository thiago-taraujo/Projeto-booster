package database

import "Projeto-booster/internal/entity"

type EventInterface interface {
	Create(event *entity.Event) error
	FindAll(page, limit int, sort string) ([]entity.Event, error)
	FindByID(id string) (*entity.Event, error)
	Update(event *entity.Event) error
	Delete(id string) error
}

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}