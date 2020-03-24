package contact

import (
	"net/http"
	"time"

	"github.com/goCrudChallenge/pkg/utl/model"
	req "github.com/goCrudChallenge/pkg/utl/model/requests"
	res "github.com/goCrudChallenge/pkg/utl/model/responses"
	"github.com/goCrudChallenge/pkg/utl/query"
	"github.com/labstack/echo"
)

var (
	// Since there's only two phone types, we'll check here
	phoneTypes        = [2]int{1, 2}
	errWrongPhoneType = echo.NewHTTPError(http.StatusInternalServerError, "Wrong phone type provided.")
)

// Create creates a contact
func (co *Contact) Create(c echo.Context, req model.Contact) (*model.Contact, error) {
	nf := 0
	if len(req.Phones) > 0 {
		for _, v := range req.Phones {
			for _, p := range phoneTypes {
				if v.PhoneTypeID != p {
					nf++
				}
			}
			if nf == 2 {
				return nil, errWrongPhoneType
			}
			nf = 0
		}
	}
	req.CreatedAt = time.Now()
	return co.cdb.Create(co.db, req)
}

// List returns list of contacts
func (co *Contact) List(c echo.Context, p *model.Pagination, r *req.ByLocation) ([]model.Contact, error) {
	q, err := query.List(r)
	if err != nil {
		return nil, err
	}
	return co.cdb.List(co.db, q, p)
}

// View views a contact's data
func (co *Contact) View(c echo.Context, id uint) (*res.ContactResponse, error) {
	return co.cdb.View(co.db, id)
}

// ByMail searches a contacat by mail
func (co *Contact) ByMail(c echo.Context, mail string, p *model.Pagination) (*[]model.Contact, error) {
	return co.cdb.ByMail(co.db, mail, p)
}

// ByPhone searches a contacat by phone
func (co *Contact) ByPhone(c echo.Context, phone *req.ByPhone) (*res.ContactResponse, error) {
	p := &model.Phone{
		Prefix: phone.Prefix,
		Number: phone.Number,
	}
	id, err := co.cdb.ByPhone(co.db, p)
	if err != nil {
		return nil, err
	}
	return co.cdb.View(co.db, id)
}

// Update struct used for passing the updated data
type Update struct {
	ID           uint
	Name         string
	CompanyID    int
	ProfileImage string
	Email        string
	BirthDate    time.Time
	StreetName   string
	StreetNumber int
	CityID       int
}

// Update for updating the data
func (co *Contact) Update(c echo.Context, r *Update) (*res.ContactResponse, error) {

	err := co.cdb.Update(co.db, &model.Contact{
		Model:        model.Model{ID: r.ID},
		Name:         r.Name,
		CompanyID:    r.CompanyID,
		ProfileImage: r.ProfileImage,
		Email:        r.Email,
		BirthDate:    r.BirthDate,
		StreetName:   r.StreetName,
		StreetNumber: r.StreetNumber,
		CityID:       r.CityID,
	})

	if err != nil {
		return nil, err
	}

	return co.cdb.View(co.db, r.ID)
}

// Delete deletes a contact
func (co *Contact) Delete(c echo.Context, id uint) error {
	contactResp, err := co.cdb.View(co.db, id)
	if err != nil {
		return err
	}

	contact := model.Contact{
		Name:         contactResp.Name,
		CompanyID:    contactResp.CompanyID,
		ProfileImage: contactResp.ProfileImage,
		Email:        contactResp.Email,
		BirthDate:    contactResp.BirthDate,
		StreetName:   contactResp.StreetName,
		StreetNumber: contactResp.StreetNumber,
		CityID:       contactResp.CityID,
	}
	contact.ID = contactResp.ID

	return co.cdb.Delete(*co.db, &contact)
}
