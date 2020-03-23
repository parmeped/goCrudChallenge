package model

// Company represents company model
type Company struct {
	Base
	Name    string  `json:"name"`
	Active  bool    `json:"active"`
	Address Address `json:"address_id"`
	Owner   Contact `json:"owner"`
}
