package main

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// IMPLICIT START OMIT

type Employee struct {
	gorm.Model
	Name      string
	CompanyID uuid.UUID
	Company   Company
}

type Company struct {
	ID   uuid.UUID
	Name string
}

// IMPLICIT END OMIT
// EXPLICIT START OMIT

type EmployeeExplicit struct {
	gorm.Model
	Name    string
	CompID  uuid.UUID
	Company Company `gorm:"foreignKey:CompID"`
}

type CompanyExplicit struct {
	ID   uuid.UUID
	Name string
}

// EXPLICIT END OMIT
// HAS-ONE START OMIT

type Building struct {
	gorm.Model
	Address Address
}

type Address struct {
	gorm.Model
	BuildingID uint
	Number     string
}

func GetAllBuildings(db *gorm.DB) ([]Building, error) {
	var buildings []Building
	err := db.Model(&Building{}).Preload("Address").Find(&buildings).Error
	return buildings, err
}

// HAS-ONE END OMIT
// HAS-MANY START OMIT

type OfficeBuilding struct {
	gorm.Model
	Offices []Office
}

type Office struct {
	gorm.Model
	BuildingID uint
	Number     string
}

func GetAllOfficeBuildings(db *gorm.DB) ([]OfficeBuilding, error) {
	var buildings []OfficeBuilding
	err := db.Model(&OfficeBuilding{}).Preload("Offices").Find(&buildings).Error
	return buildings, err
}

// HAS-MANY END OMIT
// MANY-TO-MANY START OMIT

type Developer struct {
	gorm.Model
	Name  string
	Skill []Skill `gorm:"many2many:developer_skill"` // join table
}

type Skill struct {
	gorm.Model
	Description string
	Developers  []Developer `gorm:"many2many:developer_skill"` // join table
}

func GetAllDevelopers(db *gorm.DB) ([]Developer, error) {
	var developers []Developer
	err := db.Model(&Developer{}).Preload("Skills").Find(&developers).Error
	return developers, err
}

func GetAllSkills(db *gorm.DB) ([]Skill, error) {
	var skills []Skill
	err := db.Model(&Skill{}).Preload("Developers").Find(&skills).Error
	return skills, err
}

// MANY-TO-MANY END OMIT
