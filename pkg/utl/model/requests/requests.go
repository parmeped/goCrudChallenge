package requests

type CreateReq struct {
	Name         string `json:"name" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	ProfileImage string `json:"profile_image" validate:"required"`
	BirthDate    string `json:"birth_date" validate:"required"`

	CompanyID    int    `json:"company_id" validate:"required"`
	StreetName   string `json:"street_name" validate:"required"`
	StreetNumber int    `json:"street_number" validate:"required"`
	CityID       int    `json:"city_id" validate:"required"`
	StateID      int    `json:"state_id" validate:"required"`
}

type UpdateReq struct {
	ID           int    `json:"-"`
	Name         string `json:"name,omitempty" validate:"omitempty,min=2"`
	CompanyID    int    `json:"company_id,omitempty" validate:"omitempty"`
	ProfileImage string `json:"profile_image,omitempty" validate:"omitempty"`
	Email        string `json:"email,omitempty" validate:"omitempty"`
	BirthDate    string `json:"birth_date,omitempty" validate:"omitempty"`
	StreetName   string `json:"street_name,omitempty" validate:"omitempty"`
	StreetNumber string `json:"street_number,omitempty" validate:"omitempty,min=1"`
	CityID       int    `json:"city_id,omitempty" validate:"omitempty"`
}

type ByPhone struct {
	Prefix int `json:"prefix,omitempty" validate:"omitempty,min=2"`
	Number int `json:"number,omitempty" validate:"omitempty,min=8"`
}

type ByLocation struct {
	Location string `json:"location,omitempty" validate:"omitempty,min=4"`
	ID       int    `json:"id,omitempty" validate:"omitempty,min=1"`
}
