package main

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// START OMIT

type Event struct {
	ID uuid.UUID
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.New()
	if !u.IsValid() {
		err = errors.New("can't save invalid data")
	}

	return
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	if u.ID == 1 {
		tx.Model(u).Update("role", "admin")
	}

	return
}

// END OMIT
