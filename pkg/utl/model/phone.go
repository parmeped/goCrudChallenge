package model

// TODO: add desc
type Phone struct {
	Base
	Prefix int     `json:"prefix"`
	Number int     `json:"number"`
	TypeID int     `json:"type_id"`
	Owner  Contact `json:"contact_id"`
}
