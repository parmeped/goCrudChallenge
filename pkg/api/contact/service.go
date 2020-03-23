package contact

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/goCrudChallenge/pkg/api/contact/platform/pgsql"
	"github.com/goCrudChallenge/pkg/utl/model"
	"github.com/labstack/echo"
)

// Service represents user application interface
type Service interface {
	Create(echo.Context, model.Contact) (*model.Contact, error)
	// List(echo.Context, *model.Pagination) ([]model.Contact, error)
	// View(echo.Context, int) (*model.Contact, error)
	// Delete(echo.Context, int) error
	// Update(echo.Context, *Update) (*model.Contact, error)
}

// New creates new user application service
// TODO: add interface
func New(db *pg.DB, cdb CDB) *Contact {
	return &Contact{db: db, cdb: cdb}
}

// Initialize initalizes User application service with defaults
// TODO: here the implementation gets passed to the service
func Initialize(db *pg.DB) *Contact {
	return New(db, pgsql.NewContact())
}

// User represents user application service
type Contact struct {
	db  *pg.DB
	cdb CDB
}

type CDB interface {
	Create(orm.DB, model.Contact) (*model.Contact, error)
	//View(orm.DB, int) (*model.Contact, error)
	//List(orm.DB, *gorsk.ListQuery, *gorsk.Pagination) ([]model.Contact, error)
	//Update(orm.DB, *model.Contact) error
	//Delete(orm.DB, *model.Contact) error
}

// // RBAC represents role-based-access-control interface
// type RBAC interface {
// 	User(echo.Context) *gorsk.AuthUser
// 	EnforceUser(echo.Context, int) error
// 	AccountCreate(echo.Context, gorsk.AccessRole, int, int) error
// 	IsLowerRole(echo.Context, gorsk.AccessRole) error
// }
