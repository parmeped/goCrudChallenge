package contact

import (
	"github.com/go-pg/pg"
	"github.com/goCrudChallenge/pkg/utl/model"
	"github.com/labstack/echo"
)

// Service represents user application interface
type Service interface {
	Create(echo.Context, model.Contact) (*model.Contact, error)
	List(echo.Context, *model.Pagination) ([]model.Contact, error)
	View(echo.Context, int) (*model.Contact, error)
	Delete(echo.Context, int) error
	Update(echo.Context, *Update) (*model.Contact, error)
}

// New creates new user application service
// TODO: add interface
func New(db *pg.DB) *Contact {
	return &Contact{db: db}
}

// Initialize initalizes User application service with defaults
// TODO: here the implementation gets passed to the service
func Initialize(db *pg.DB) *Contact {
	return New(db)
}

// User represents user application service
type Contact struct {
	db *pg.DB
}

// // UDB represents user repository interface
// type UDB interface {
// 	Create(orm.DB, gorsk.User) (*gorsk.User, error)
// 	View(orm.DB, int) (*gorsk.User, error)
// 	List(orm.DB, *gorsk.ListQuery, *gorsk.Pagination) ([]gorsk.User, error)
// 	Update(orm.DB, *gorsk.User) error
// 	Delete(orm.DB, *gorsk.User) error
// }

// // RBAC represents role-based-access-control interface
// type RBAC interface {
// 	User(echo.Context) *gorsk.AuthUser
// 	EnforceUser(echo.Context, int) error
// 	AccountCreate(echo.Context, gorsk.AccessRole, int, int) error
// 	IsLowerRole(echo.Context, gorsk.AccessRole) error
// }
