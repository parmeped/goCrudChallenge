package model

// TODO: add desc
type City struct {
	Base
	Name      string
	StateID   int `gorm:"index"`
	Companies []Company
	Contacts  []Contact
}
