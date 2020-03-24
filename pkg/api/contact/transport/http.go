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

// HTTP represents contact http service
type HTTP struct {
	svc contact.Service
}

// NewHTTP creates new contact http service
func NewHTTP(svc contact.Service, er *echo.Group) {
	h := HTTP{svc}
	cg := er.Group("/contacts")

	cg.POST("", h.create)

	cg.GET("/listByMail/:mail", h.byMail)

	cg.GET("/byPhone", h.byPhone)

	cg.GET("/listByLocation/:searchParam/:id", h.list)

	cg.GET("/:id", h.view)

	cg.PATCH("/:id", h.update)

	cg.DELETE("/:id", h.delete)

	cg.Any("/", h.notFound)
}

func (h *HTTP) create(c echo.Context) error {
	r := new(req.CreateReq)

	if err := c.Bind(r); err != nil {
		return err
	}

	birthD, err := time.Parse(time.RFC3339, r.BirthDate)
	if err != nil {
		return model.ErrParsingDate
	}

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

// internal struct
type listResponse struct {
	Contacts []model.Contact `json:"contacts"`
	Page     int             `json:"page"`
}

// list all contacts given a location search param
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

// view a certain contact given its ID
func (h *HTTP) view(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return model.ErrBadRequest
	}
	uid := uint(id)
	result, err := h.svc.View(c, uid)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

// Search a contact by mail
func (h *HTTP) byMail(c echo.Context) error {
	mail := c.Param("mail")
	p := new(model.PaginationReq)
	if err := c.Bind(p); err != nil {
		return err
	}

	result, err := h.svc.ByMail(c, mail, p.Transform())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

// search by phone number
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

// Update a contact's data
func (h *HTTP) update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return model.ErrParsingID
	}
	uid := uint(id)

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
		ID:           uid,
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

// delete a contact through its ID
func (h *HTTP) delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return model.ErrBadRequest
	}

	uid := uint(id)
	if err := h.svc.Delete(c, uid); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

// should be called when no other endpoint catches the request
func (h *HTTP) notFound(c echo.Context) error {
	return c.NoContent(http.StatusNotFound)
}
