package model

import "time"

// Contact represents contact model
type Contact struct {
	Base
	Name         string    `json:"name"`
	Active       bool      `json:"active"`
	CompanyID    int       `json:"company_id"`
	ProfileImage string    `json:"profile_image"` // TODO: this should be a BLOB, or something similar. See how to save this data.
	Email        string    `json:"email"`
	BirthDate    time.Time `json:"birth_date"`
	Address      Address   `json:"address_id"`
}
