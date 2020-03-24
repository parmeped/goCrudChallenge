package model

import "strconv"

// TODO: add desc
type Phone struct {
	Base
	Prefix int     `json:"prefix"`
	Number int     `json:"number"`
	TypeID int     `json:"type_id"`
	Owner  Contact `json:"contact_id"`
}

func (p *Phone) GetPhone() string {
	return strconv.Itoa(p.Prefix) + "-" + strconv.Itoa(p.Number)
}
