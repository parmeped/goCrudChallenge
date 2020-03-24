package model

// Phone model
type Phone struct {
	Model
	Prefix      int  `json:"prefix" validate:"required"`
	Number      int  `json:"number" validate:"required"`
	PhoneTypeID int  `gorm:"index" json:"type_id"`
	ContactID   uint `gorm:"index`
}
