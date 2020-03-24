package model

// Company represents company model
type Company struct {
	Base
	Name         string `json:"name"`
	StreetName   string `json:"street_name"`
	StreetNumber int    `json:"street_number"`
	CityID       int    `json:"city_id"`
}
