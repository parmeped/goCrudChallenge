package model

// Company model
type Company struct {
	Model
	Name         string
	StreetName   string
	StreetNumber int
	CityID       int `gorm:"index"`
	Contacts     []Contact
}
