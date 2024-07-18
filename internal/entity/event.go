package entity

import (
	"errors"
	"time"

)

var (
	ErrIDIsRequired          = errors.New("ID is required")
	ErrInvalidID             = errors.New("ID is invalid")
	ErrNameIsRequired        = errors.New("name is required")
	ErrDescriptionIsRequired = errors.New("description is required")
	ErrStartDateIsRequired   = errors.New("start date is required")
	ErrFinishDateIsRequired  = errors.New("finish is required")
	ErrInvalidDate           = errors.New("invalid start or finish date")
)

type Event struct {
	ID          entity.ID `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	FinishDate  time.Time `json:"finish_date"`
}

func NewEvent(name, description string, startDate, finishDate time.Time) (*Event, error) {
	event := &Event{
		ID:          entity.NewID(),
		Name:        name,
		Description: description,
		StartDate:   startDate,
		FinishDate:  finishDate,
	}
	if err := event.IsValid(); err != nil {
		return nil, err
	}
	return event, nil
}

func (e *Event) IsValid() error {
	if e.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID(e.ID.String()); err != nil {
		return ErrInvalidID
	}
	if e.Name == "" {
		return ErrNameIsRequired
	}
	if e.Description == "" {
		return ErrDescriptionIsRequired
	}
	if (e.StartDate.GoString() == "") {
		return ErrStartDateIsRequired
	}
	if (e.FinishDate.GoString() == "") {
		return ErrFinishDateIsRequired
	}
	if e.StartDate.After(e.FinishDate) {
		return ErrInvalidDate
	}
	return nil
}
