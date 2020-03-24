package model

// TODO: add desc
type State struct {
	Base
	Name   string `json:"name"`
	Cities []City
}
