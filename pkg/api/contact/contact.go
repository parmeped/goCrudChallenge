package contact

import (
	"time"

	"github.com/goCrudChallenge/pkg/utl/model"
	res "github.com/goCrudChallenge/pkg/utl/model/responses"
	"github.com/labstack/echo"
)

//TODO: Inteface proper implementation

//"github.com/ribice/gorsk/pkg/utl/query"

// TODO: here's where bussiness logic would be.
func (co *Contact) Create(c echo.Context, req model.Contact) (*model.Contact, error) {
	req.CreatedAt = time.Now()
	return co.cdb.Create(co.db, req)
}

// // List returns list of users
// func (u *User) List(c echo.Context, p *gorsk.Pagination) ([]gorsk.User, error) {
// 	au := u.rbac.User(c)
// 	q, err := query.List(au)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return u.udb.List(u.db, q, p)
// }

// TODO: add desc
func (co *Contact) View(c echo.Context, id int) (*res.ContactResponse, error) {
	return co.cdb.View(co.db, id)
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
