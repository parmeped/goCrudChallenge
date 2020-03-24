package model

// TODO: add desc
type City struct {
	Base
	Name    string `json:"name"`
	StateID int    `json:"state_id"`
}
