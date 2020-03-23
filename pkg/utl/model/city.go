package model

// TODO: add desc
type City struct {
	Base
	Name  string `json:"name"`
	State State  `json:"state_id"`
}
