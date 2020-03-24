package model

import (
	"errors"

	"github.com/labstack/echo"
)

var (
	// ErrBadRequest (400) is returned for bad request (validation)
	ErrBadRequest = echo.NewHTTPError(400)
	// ErrReqWithStreetNumber error parsing street number
	ErrReqWithStreetNumber = errors.New("Error parsing the street number")
	// ErrParsingID error parsing the id
	ErrParsingID = errors.New("Error parsing the id")
	// ErrParsingDate error parsing date
	ErrParsingDate = errors.New("Error parsing the date")
	// ErrWrongSearchParameters wrong parameters provided
	ErrWrongSearchParameters = errors.New("wrong parameters. Accepted parameters are company or city")
)
