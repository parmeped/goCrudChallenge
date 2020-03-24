package model

// State model
type State struct {
	Model
	Name   string `json:"name"`
	Cities []City
}
