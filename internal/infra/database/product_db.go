package database

import (
	"Projeto-booster/internal/entity"
	"gorm.io/gorm"
)

type Event struct {
	DB *gorm.DB
}

func NewEvent(db *gorm.DB) *Event {
	return &Event{DB: db}
}

func (e *Event) Create(event *entity.Event) error {
	return e.DB.Create(event).Error
}

func (e *Event) FindByID(id string) (*entity.Event, error) {
	var event entity.Event
	err := e.DB.First(&event, "id = ?", id).Error
	return &event, err
}

func (e *Event) FindAll(page, limit int, sort string) ([]entity.Event, error) {
	var events []entity.Event
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		err = e.DB.Limit(limit).Offset((page - 1) * limit).Order("start_date " + sort).Find(&events).Error
	} else {
		err = e.DB.Order("start_date " + sort).Find(&events).Error
	}
	return events, err
}

func (e *Event) Update(event *entity.Event) error {
	_, err := e.FindByID(event.ID.String())
	if err != nil {
		return err
	}
	return e.DB.Save(event).Error
}

func (e *Event) Delete(id string) error {
	event, err := e.FindByID(id)
	if err != nil {
		return err
	}
	return e.DB.Delete(event).Error
}
