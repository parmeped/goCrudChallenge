package model

import "time"

// Contact model
type Contact struct {
	Model
	Name         string
	CompanyID    int    `gorm:"index"`
	ProfileImage string // TODO: [IMPROVEMENT]: This should be a BLOB instead of a simple url.
	Email        string
	BirthDate    time.Time
	StreetName   string
	StreetNumber int
	CityID       int `gorm:"foreignkey:city_id"`
	Phones       []Phone
}
