package model

// City model
type City struct {
	Model
	Name      string
	StateID   int `gorm:"index"`
	Companies []Company
	Contacts  []Contact
}
