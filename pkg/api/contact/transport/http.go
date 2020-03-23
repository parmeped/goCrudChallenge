package transport

import (
	"net/http"
	"time"

	"github.com/goCrudChallenge/pkg/api/contact"
	"github.com/goCrudChallenge/pkg/utl/model"

	"github.com/labstack/echo"
)

// HTTP represents user http service
type HTTP struct {
	svc contact.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc contact.Service, er *echo.Group) {
	h := HTTP{svc}
	ur := er.Group("/contacts")
	// swagger:route POST /v1/users users userCreate
	// Creates new user account.
	// responses:
	//  200: userResp
	//  400: errMsg
	//  401: err
	//  403: errMsg
	//  500: err
	ur.POST("", h.create)

	// swagger:operation GET /v1/users users listUsers
	// ---
	// summary: Returns list of users.
	// description: Returns list of users. Depending on the user role requesting it, it may return all users for SuperAdmin/Admin users, all company/location users for Company/Location admins, and an error for non-admin users.
	// parameters:
	// - name: limit
	//   in: query
	//   description: number of results
	//   type: int
	//   required: false
	// - name: page
	//   in: query
	//   description: page number
	//   type: int
	//   required: false
	// responses:
	//   "200":
	//     "$ref": "#/responses/userListResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	//ur.GET("", h.list)

	// swagger:operation GET /v1/users/{id} users getUser
	// ---
	// summary: Returns a single user.
	// description: Returns a single user by its ID.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of user
	//   type: int
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/userResp"
	//   "400":
	//     "$ref": "#/responses/err"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "404":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	//ur.GET("/:id", h.view)

	// swagger:operation PATCH /v1/users/{id} users userUpdate
	// ---
	// summary: Updates user's contact information
	// description: Updates user's contact information -> first name, last name, mobile, phone, address.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of user
	//   type: int
	//   required: true
	// - name: request
	//   in: body
	//   description: Request body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/userUpdate"
	// responses:
	//   "200":
	//     "$ref": "#/responses/userResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	//ur.PATCH("/:id", h.update)

	// swagger:operation DELETE /v1/users/{id} users userDelete
	// ---
	// summary: Deletes a user
	// description: Deletes a user with requested ID.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of user
	//   type: int
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/ok"
	//   "400":
	//     "$ref": "#/responses/err"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	//ur.DELETE("/:id", h.delete)
}

// // Custom errors
// var (
// 	ErrPasswordsNotMaching = echo.NewHTTPError(http.StatusBadRequest, "passwords do not match")
// )

// User create request
// swagger:model userCreate
type createReq struct {
	Name         string    `json:"first_name" validate:"required"`
	Email        string    `json:"email" validate:"required,email"`
	ProfileImage string    `json:"profile_image" validate:"required"`
	BirthDate    time.Time `json:"birth_date" validate:"required"`

	CompanyID    int    `json:"company_id" validate:"required"`
	StreetName   string `json:"street_name" validate:"required"`
	StreetNumber int    `json:"street_number" validate:"required"`
	CityID       int    `json:"city_id" validate:"required"`
	StateID      int    `json:"state_id" validate:"required"`
}

func (h *HTTP) create(c echo.Context) error {
	r := new(createReq)

	if err := c.Bind(r); err != nil {

		return err
	}

	// TODO: validate if address exists. should go on the address service create?

	cnt, err := h.svc.Create(c, model.Contact{
		Name:         r.Name,
		Active:       true,
		CompanyID:    r.CompanyID,
		ProfileImage: r.ProfileImage,
		Email:        r.Email,
		BirthDate:    r.BirthDate,
		StreetName:   r.StreetName,
		StreetNumber: r.StreetNumber,
		CityID:       r.CityID,
		StateID:      r.StateID,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, cnt)
}

// type listResponse struct {
// 	Users []gorsk.User `json:"users"`
// 	Page  int          `json:"page"`
// }

// func (h *HTTP) list(c echo.Context) error {
// 	p := new(gorsk.PaginationReq)
// 	if err := c.Bind(p); err != nil {
// 		return err
// 	}

// 	result, err := h.svc.List(c, p.Transform())

// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(http.StatusOK, listResponse{result, p.Page})
// }

// func (h *HTTP) view(c echo.Context) error {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return gorsk.ErrBadRequest
// 	}

// 	result, err := h.svc.View(c, id)
// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(http.StatusOK, result)
// }

// // User update request
// // swagger:model userUpdate
// type updateReq struct {
// 	ID        int    `json:"-"`
// 	FirstName string `json:"first_name,omitempty" validate:"omitempty,min=2"`
// 	LastName  string `json:"last_name,omitempty" validate:"omitempty,min=2"`
// 	Mobile    string `json:"mobile,omitempty"`
// 	Phone     string `json:"phone,omitempty"`
// 	Address   string `json:"address,omitempty"`
// }

// func (h *HTTP) update(c echo.Context) error {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return gorsk.ErrBadRequest
// 	}

// 	req := new(updateReq)
// 	if err := c.Bind(req); err != nil {
// 		return err
// 	}

// 	usr, err := h.svc.Update(c, &user.Update{
// 		ID:        id,
// 		FirstName: req.FirstName,
// 		LastName:  req.LastName,
// 		Mobile:    req.Mobile,
// 		Phone:     req.Phone,
// 		Address:   req.Address,
// 	})

// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(http.StatusOK, usr)
// }

// func (h *HTTP) delete(c echo.Context) error {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return gorsk.ErrBadRequest
// 	}

// 	if err := h.svc.Delete(c, id); err != nil {
// 		return err
// 	}

// 	return c.NoContent(http.StatusOK)
// }
