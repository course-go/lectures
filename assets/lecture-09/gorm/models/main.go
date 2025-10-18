package main

import (
	"time"

	"gorm.io/gorm"
)

// MODEL START OMIT

type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// MODEL END OMIT
// USER START OMIT

type User struct {
	ID   string
	Name string
}

// USER END OMIT
// TABLENAME START OMIT

type Tabler interface {
	TableName() string
}

func (User) TableName() string {
	return "profiles"
}

// TABLENAME END OMIT
// ANIMAL START OMIT

type Animal struct {
	ID   int64
	UUID string `gorm:"primaryKey"`
	Name string
	Age  int64
}

// ANIMAL END OMIT
// CATEGORY START OMIT

type Category struct {
	ID        uint      // `id`
	Name      string    // `name`
	CreatedAt time.Time // `created_at`
}

// CATEGORY END OMIT
// SHOP START OMIT

type Shop struct {
	ID       string `gorm:"column:shop_id"`
	Location string
}

// SHOP END OMIT
// ORDER START OMIT

type Order struct {
	CreatedAt time.Time `gorm:"autoCreateTime:false"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:false"`
}

// ORDER END OMIT
