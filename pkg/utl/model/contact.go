package model

import "time"

// Contact represents contact model
type Contact struct {
	Base
	Name         string    `json:"name"`
	CompanyID    int       `json:"company_id"`
	ProfileImage string    `json:"profile_image"` // TODO: this should be a BLOB, or something similar. See how to save this data.
	Email        string    `json:"email"`
	BirthDate    time.Time `json:"birth_date"`
	StreetName   string    `json:"street_name"`
	StreetNumber int       `json:"street_number"`
	CityID       int       `json:"city_id"`
	Phones       []Phone   `json:"phones"`
}
