package responses

import (
	"time"

	"github.com/goCrudChallenge/pkg/utl/model"
)

// ContactResponse is returned when viewing a contact's data
type ContactResponse struct {
	ID           uint          `json:"id"`
	Name         string        `json:"name"`
	CompanyID    int           `json:"company_id"`
	CompanyName  string        `json:"company_name"`
	ProfileImage string        `json:"profile_image"`
	Email        string        `json:"email"`
	BirthDate    time.Time     `json:"birth_date"`
	StreetName   string        `json:"street_name"`
	StreetNumber int           `json:"street_number"`
	CityID       int           `json:"city_id"`
	CityName     string        `json:"city_name"`
	StateName    string        `json:"state_name"`
	Phones       []model.Phone `json:"phones"`
}
