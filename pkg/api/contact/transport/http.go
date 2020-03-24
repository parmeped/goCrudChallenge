package transport

import (
	"net/http"
	"strconv"
	"time"

	"github.com/goCrudChallenge/pkg/api/contact"
	"github.com/goCrudChallenge/pkg/utl/model"
	req "github.com/goCrudChallenge/pkg/utl/model/requests"

	"github.com/labstack/echo"
)

// HTTP represents user http service
type HTTP struct {
	svc contact.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc contact.Service, er *echo.Group) {
	h := HTTP{svc}
	cg := er.Group("/contacts")

	cg.POST("", h.create)

	cg.GET("/listByMail/:mail", h.byMail)

	cg.GET("/listByPhone", h.byPhone)

	cg.GET("/listByLocation/:searchParam/:id", h.list)

	cg.GET("/:id", h.view)

	cg.PATCH("/:id", h.update)

	cg.DELETE("/:id", h.delete)
}

// TODO: Updates are not logging updated time

func (h *HTTP) create(c echo.Context) error {
	r := new(req.CreateReq)

	if err := c.Bind(r); err != nil {

		return err
	}

	// TODO: check this time parsing
	birthD, err := time.Parse(time.RFC3339, r.BirthDate)
	if err != nil {
		return model.ErrParsingDate
	}

	if err != nil {
		return err
	}

	// TODO: phone missing, remove Active
	cnt, err := h.svc.Create(c, model.Contact{
		Name:         r.Name,
		CompanyID:    r.CompanyID,
		ProfileImage: r.ProfileImage,
		Email:        r.Email,
		BirthDate:    birthD,
		StreetName:   r.StreetName,
		StreetNumber: r.StreetNumber,
		CityID:       r.CityID,
		Phones:       r.Phones,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, cnt)
}

type listResponse struct {
	Contacts []model.Contact `json:"contacts"`
	Page     int             `json:"page"`
}

func (h *HTTP) list(c echo.Context) error {
	p := new(model.PaginationReq)
	if err := c.Bind(p); err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return model.ErrBadRequest
	}

	byLocationReq := &req.ByLocation{
		c.Param("searchParam"),
		id,
	}

	result, err := h.svc.List(c, p.Transform(), byLocationReq)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, listResponse{result, p.Page})
}

func (h *HTTP) view(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return model.ErrBadRequest
	}

	result, err := h.svc.View(c, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

// TODO: add pagination
func (h *HTTP) byMail(c echo.Context) error {
	mail := c.Param("mail")

	result, err := h.svc.ByMail(c, mail)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (h *HTTP) byPhone(c echo.Context) error {
	req := &req.ByPhone{}
	if err := c.Bind(req); err != nil {
		return err
	}

	result, err := h.svc.ByPhone(c, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (h *HTTP) update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return model.ErrParsingId
	}

	req := new(req.UpdateReq)
	if err := c.Bind(req); err != nil {
		return err
	}

	streetNum, err := strconv.Atoi(req.StreetNumber)
	if err != nil {
		return model.ErrReqWithStreetNumber
	}

	birthD, err := time.Parse(time.RFC3339, req.BirthDate)
	if err != nil {
		return model.ErrParsingDate
	}

	contact, err := h.svc.Update(c, &contact.Update{
		ID:           id,
		Name:         req.Name,
		CompanyID:    req.CompanyID,
		ProfileImage: req.ProfileImage,
		Email:        req.Email,
		BirthDate:    birthD,
		StreetName:   req.StreetName,
		StreetNumber: streetNum,
		CityID:       req.CityID,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, contact)
}

func (h *HTTP) delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return model.ErrBadRequest
	}

	if err := h.svc.Delete(c, id); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
