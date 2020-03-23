package model

import "strconv"

// TODO: Remove
type Address struct {
	Base
	Active  bool   `json:"active"`
	Street  string `json:"street_name"`
	Number  int    `json:"street_number"`
	CityID  int    `json:"city_id"`
	StateID int    `json:"state_id"`
}

func (a *Address) GetAddress() string {
	return "Street " + a.Street + " " + strconv.FormatInt(int64(a.Number), 10)
}
