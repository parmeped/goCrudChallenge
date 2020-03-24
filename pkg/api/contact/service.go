package contact

import (
	"github.com/goCrudChallenge/pkg/api/contact/platform/pgsql"
	"github.com/goCrudChallenge/pkg/utl/model"
	req "github.com/goCrudChallenge/pkg/utl/model/requests"
	res "github.com/goCrudChallenge/pkg/utl/model/responses"
	"github.com/goCrudChallenge/pkg/utl/query"
	"github.com/jinzhu/gorm"

	"github.com/labstack/echo"
)

// Service represents contact application interface
type Service interface {
	Create(echo.Context, model.Contact) (*model.Contact, error)
	List(echo.Context, *model.Pagination, *req.ByLocation) ([]model.Contact, error)
	View(echo.Context, uint) (*res.ContactResponse, error)
	Delete(echo.Context, uint) error
	Update(echo.Context, *Update) (*res.ContactResponse, error)
	ByMail(echo.Context, string, *model.Pagination) (*[]model.Contact, error)
	ByPhone(echo.Context, *req.ByPhone) (*res.ContactResponse, error)
}

// New creates new contact application service
func New(db *gorm.DB, cdb CDB) *Contact {
	return &Contact{db: db, cdb: cdb}
}

// Initialize initalizes contact application service with defaults
func Initialize(db *gorm.DB) *Contact {
	return New(db, pgsql.NewContact())
}

// Contact represents the contacts application service
type Contact struct {
	db  *gorm.DB
	cdb CDB
}

// CDB represents the contact DB interface
type CDB interface {
	Create(*gorm.DB, model.Contact) (*model.Contact, error)
	View(*gorm.DB, uint) (*res.ContactResponse, error)
	List(*gorm.DB, *query.ListQuery, *model.Pagination) ([]model.Contact, error)
	Delete(gorm.DB, *model.Contact) error
	Update(*gorm.DB, *model.Contact) error
	ByMail(*gorm.DB, string, *model.Pagination) (*[]model.Contact, error)
	ByPhone(*gorm.DB, *model.Phone) (uint, error)
}
