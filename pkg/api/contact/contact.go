package contact

import (
	"time"

	"github.com/goCrudChallenge/pkg/utl/model"
	req "github.com/goCrudChallenge/pkg/utl/model/requests"
	res "github.com/goCrudChallenge/pkg/utl/model/responses"
	"github.com/goCrudChallenge/pkg/utl/query"
	"github.com/labstack/echo"
)

//TODO: Inteface proper implementation

//"github.com/ribice/gorsk/pkg/utl/query"

// TODO: here's where bussiness logic would be.
func (co *Contact) Create(c echo.Context, req model.Contact) (*model.Contact, error) {
	req.CreatedAt = time.Now()
	return co.cdb.Create(co.db, req)
}

// List returns list of contacts
func (co *Contact) List(c echo.Context, p *model.Pagination, byLocReq *req.ByLocation) ([]model.Contact, error) {
	q, err := query.List(byLocReq)
	if err != nil {
		return nil, err
	}
	return co.cdb.List(co.db, q, p)
}

// TODO: add desc
func (co *Contact) View(c echo.Context, id int) (*res.ContactResponse, error) {
	return co.cdb.View(co.db, id)
}

// TODO: add desc
func (co *Contact) ByMail(c echo.Context, mail string) (*res.ContactResponse, error) {
	return co.cdb.ByMail(co.db, mail)
}

// TODO: add desc
func (co *Contact) ByPhone(c echo.Context, phone *req.ByPhone) (*res.ContactResponse, error) {
	return co.cdb.ByPhone(co.db, phone)
}

// Delete deletes a user
func (co *Contact) Delete(c echo.Context, id int) error {
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

	return co.cdb.Delete(co.db, &contact)
}

// Update contains user's information used for updating
type Update struct {
	ID           int
	Name         string
	CompanyID    int
	ProfileImage string
	Email        string
	BirthDate    time.Time
	StreetName   string
	StreetNumber int
	CityID       int
}

// Update updates user's contact information
func (co *Contact) Update(c echo.Context, r *Update) (*res.ContactResponse, error) {

	if err := co.cdb.Update(co.db, &model.Contact{
		Base:         model.Base{ID: r.ID},
		Name:         r.Name,
		CompanyID:    r.CompanyID,
		ProfileImage: r.ProfileImage,
		Email:        r.Email,
		BirthDate:    r.BirthDate,
		StreetName:   r.StreetName,
		StreetNumber: r.StreetNumber,
		CityID:       r.CityID,
	}); err != nil {
		return nil, err
	}

	return co.cdb.View(co.db, r.ID)
}
