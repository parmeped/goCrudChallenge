package pgsql

import (
	"net/http"
	"strings"

	"github.com/goCrudChallenge/pkg/utl/model"
	res "github.com/goCrudChallenge/pkg/utl/model/responses"
	"github.com/goCrudChallenge/pkg/utl/query"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// NewContact returns a new contact for later calls
func NewContact() *Contact {
	return &Contact{}
}

// Contact struct
type Contact struct{}

// Custom errors
var (
	ErrAlreadyExists      = echo.NewHTTPError(http.StatusBadRequest, "Username or email already exists.")
	ErrCityDoesntExist    = echo.NewHTTPError(http.StatusBadRequest, "City  doesn't exist.")
	ErrCompanyDoesntExist = echo.NewHTTPError(http.StatusBadRequest, "Company doesn't exist.")
	ErrNotExists          = echo.NewHTTPError(http.StatusInternalServerError, "ID wasn't found.")
	ErrDeleteFailed       = echo.NewHTTPError(http.StatusInternalServerError, "Delete failed.")
	ErrNotFound           = echo.NewHTTPError(http.StatusInternalServerError, "The contact wasn't found.")
)

// Create creates a new contact on database
func (co *Contact) Create(db *gorm.DB, cont model.Contact) (*model.Contact, error) {
	contact := model.Contact{}
	city := model.City{}
	company := model.Company{}

	db.Select("ID").Where("lower(name) = ? or lower(email) = ? and deleted_at is null", strings.ToLower(cont.Name), strings.ToLower(cont.Email)).First(&contact)
	db.Select("ID").Where("id = ?", cont.CityID).First(&city)
	db.Select("ID").Where("id = ?", cont.CompanyID).First(&company)

	if contact.ID != 0 {
		return nil, ErrAlreadyExists
	}

	if city.ID == 0 {
		return nil, ErrCityDoesntExist
	}

	if company.ID == 0 {
		return nil, ErrCompanyDoesntExist
	}

	db.Create(&cont)

	return &cont, nil
}

// View returns a contact by ID
func (co *Contact) View(db *gorm.DB, id uint) (*res.ContactResponse, error) {
	var r = new(res.ContactResponse)
	contact := model.Contact{}
	city := model.City{}
	company := model.Company{}
	phones := []model.Phone{}
	state := model.State{}

	// TODO: [IMPROVEMENT] Make a single query to improve performance and reading
	db.Where("id = ?", id).Find(&contact)
	db.Where("id = ?", contact.CityID).Find(&city)
	db.Where("id = ?", contact.CompanyID).Find(&company)
	db.Where("id = ?", city.StateID).Find(&state)
	db.Where("contact_id = ?", id).Find(&phones)

	if contact.ID == 0 {
		return nil, ErrNotExists
	}

	r.ID = contact.ID
	r.CityID = contact.CityID
	r.CityName = city.Name
	r.CompanyID = contact.CompanyID
	r.CompanyName = company.Name
	r.Email = contact.Email
	r.ProfileImage = contact.ProfileImage
	r.StreetName = contact.StreetName
	r.StreetNumber = contact.StreetNumber
	r.Name = contact.Name
	r.StateName = state.Name
	r.Phones = phones

	return r, nil
}

// ByMail searches contacts by mail
func (co *Contact) ByMail(db *gorm.DB, mail string, p *model.Pagination) (*[]model.Contact, error) {
	contacts := &[]model.Contact{}
	phones := []model.Phone{}
	if err := db.Where("UPPER(email) LIKE ?", strings.ToUpper("%"+mail+"%")).Limit(p.Limit).Offset(p.Offset).Find(&contacts).Error; err != nil {
		return nil, err
	}

	if len(*contacts) < 0 {
		return nil, ErrNotFound
	}

	for k, v := range *contacts {

		db.Where("contact_id = ?", v.ID).Find(&phones)
		(*contacts)[k].Phones = phones
	}

	return contacts, nil
}

// ByPhone searches a contact by phone
func (co *Contact) ByPhone(db *gorm.DB, phone *model.Phone) (uint, error) {

	if err := db.Where("prefix = ? AND number = ?", phone.Prefix, phone.Number).First(&phone).Error; err != nil {
		return 0, err
	}

	return phone.ContactID, nil
}

// Update the contact's info
func (co *Contact) Update(db *gorm.DB, r *model.Contact) error {
	city := model.City{}
	company := model.Company{}

	db.Select("ID").Where("id = ?", r.CityID).First(&city)
	db.Select("ID").Where("id = ?", r.CompanyID).First(&company)

	if city.ID == 0 {
		return ErrCityDoesntExist
	}

	if company.ID == 0 {
		return ErrCompanyDoesntExist
	}

	u := model.Contact{
		Model:        model.Model{ID: r.ID},
		Name:         r.Name,
		CompanyID:    r.CompanyID,
		ProfileImage: r.ProfileImage,
		Email:        r.Email,
		BirthDate:    r.BirthDate,
		StreetName:   r.StreetName,
		StreetNumber: r.StreetNumber,
		CityID:       r.CityID,
		Phones:       r.Phones,
	}

	// should retrieve from the db first, if it doesn't found, toss an error
	if err := db.Where("id = ?", r.ID).First(&r).Error; err != nil {
		return ErrNotExists
	}

	if err := db.Save(&u).Error; err != nil {
		return err
	}

	return nil
}

// List returns list of all contacts depending on location query
func (co *Contact) List(db *gorm.DB, qp *query.ListQuery, p *model.Pagination) ([]model.Contact, error) {

	contacts := []model.Contact{}
	if err := db.Where("deleted_at is null").Where(qp.Query, qp.ID).Limit(p.Limit).Offset(p.Offset).Find(&contacts).Order("contact.id desc").Error; err != nil {
		return nil, err
	}

	return contacts, nil
}

// Delete sets deleted_at for a contact (soft delete)
func (co *Contact) Delete(db gorm.DB, contact *model.Contact) error {
	phone := []model.Phone{}
	if err := db.Delete(&contact).Error; err != nil {
		return ErrDeleteFailed
	}
	db.Where("contact_id = ?", contact.ID).Delete(&phone)

	return nil
}
