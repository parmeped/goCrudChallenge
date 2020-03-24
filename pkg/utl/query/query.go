package query

import (
	"github.com/goCrudChallenge/pkg/utl/model"
	req "github.com/goCrudChallenge/pkg/utl/model/requests"
)

type ListQuery struct {
	Query string
	ID    int
}

// List prepares the list queries
func List(param *req.ByLocation) (*ListQuery, error) {
	switch true {
	case param.Location == "company":
		return &ListQuery{Query: "company_id = ?", ID: param.ID}, nil
	case param.Location == "city":
		return &ListQuery{Query: "city_id = ?", ID: param.ID}, nil
	default:
		return nil, model.ErrWrongSearchParameters
	}
}
