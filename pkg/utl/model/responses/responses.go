package responses

import (
	"time"

	"github.com/goCrudChallenge/pkg/utl/model"
)

type ContactResponse struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	Active       bool          `json:"active"`
	CompanyID    int           `json:"company_id"`
	CompanyName  string        `json:"company_name"`
	ProfileImage string        `json:"profile_image"` // TODO: this should be a BLOB, or something similar. See how to save this data.
	Email        string        `json:"email"`
	BirthDate    time.Time     `json:"birth_date"`
	StreetName   string        `json:"street_name"`
	StreetNumber int           `json:"street_number"`
	CityID       int           `json:"city_id"`
	CityName     string        `json:"city_name"`
	StateName    string        `json:"state_name"`
	Phones       []model.Phone `json:"phones"`
}
