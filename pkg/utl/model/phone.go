package model

// TODO: add desc
type Phone struct {
	Base
	Prefix      int
	Number      int
	PhoneTypeID int `gorm:"index"`
	ContactID   int `gorm:"index"`
}
