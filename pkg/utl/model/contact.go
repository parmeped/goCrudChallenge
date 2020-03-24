package model

import "time"

// Contact represents contact model
type Contact struct {
	Base
	Name         string
	CompanyID    int    `gorm:"index"`
	ProfileImage string // TODO: this should be a BLOB, or something similar. See how to save this data.
	Email        string
	BirthDate    time.Time
	StreetName   string
	StreetNumber int
	CityID       int `gorm:"index"`
	Phones       []Phone
}
