package model

// Company represents company model
type Company struct {
	Base
	Name         string
	StreetName   string
	StreetNumber int
	CityID       int `gorm:"index"`
	Contacts     []Contact
}
