package main

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// START OMIT

type Event struct {
	ID       uuid.UUID
	Time     time.Time
	Readonly bool
}

func (e *Event) BeforeCreate(tx *gorm.DB) error {
	if e.Time.Before(time.Now()) {
		return errors.New("event time is in the past")
	}
	e.ID = uuid.New()
	return nil
}

func (e *Event) BeforeUpdate(tx *gorm.DB) error {
	if e.Readonly {
		return errors.New("event is read-only")
	}
	return nil
}

// END OMIT
