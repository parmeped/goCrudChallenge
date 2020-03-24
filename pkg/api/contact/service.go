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

// Service represents user application interface
type Service interface {
	Create(echo.Context, model.Contact) (*model.Contact, error)
	List(echo.Context, *model.Pagination, *req.ByLocation) ([]model.Contact, error)
	View(echo.Context, int) (*res.ContactResponse, error)
	Delete(echo.Context, int) error
	Update(echo.Context, *Update) (*res.ContactResponse, error)
	ByMail(echo.Context, string) (*res.ContactResponse, error)
	ByPhone(echo.Context, *req.ByPhone) (*res.ContactResponse, error)
}

// New creates new user application service
// TODO: add interface
func New(db *gorm.DB, cdb CDB) *Contact {
	return &Contact{db: db, cdb: cdb}
}

// Initialize initalizes User application service with defaults
// TODO: here the implementation gets passed to the service
func Initialize(db *gorm.DB) *Contact {
	return New(db, pgsql.NewContact())
}

// User represents user application service
type Contact struct {
	db  *gorm.DB
	cdb CDB
}

type CDB interface {
	Create(gorm.DB, model.Contact) (*model.Contact, error)
	View(gorm.DB, int) (*res.ContactResponse, error)
	List(gorm.DB, *query.ListQuery, *model.Pagination) ([]model.Contact, error)
	Delete(gorm.DB, *model.Contact) error
	Update(gorm.DB, *model.Contact) error
	ByMail(gorm.DB, string) (*res.ContactResponse, error)
	ByPhone(gorm.DB, *req.ByPhone) (*res.ContactResponse, error)
}
