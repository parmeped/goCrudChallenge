package pgsql

import (
	"net/http"
	"strings"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/goCrudChallenge/pkg/utl/model"
	res "github.com/goCrudChallenge/pkg/utl/model/responses"
	"github.com/labstack/echo"
)

// TODO: add desc
func NewContact() *Contact {
	return &Contact{}
}

// TODO: add desc

type Contact struct{}

// Custom errors
var (
	ErrAlreadyExists = echo.NewHTTPError(http.StatusInternalServerError, "Username or email already exists.")
)

// Create creates a new user on database
// TODO: validate here with queries if state, company or city exists. There could also be a comment about addresses.
func (c *Contact) Create(db orm.DB, cont model.Contact) (*model.Contact, error) {
	var contact = new(model.Contact)
	err := db.Model(contact).Where("lower(name) = ? or lower(email) = ? and deleted_at is null",
		strings.ToLower(cont.Name), strings.ToLower(cont.Email)).Select()
	if (err == nil) || (err != nil && err != pg.ErrNoRows) {
		return nil, ErrAlreadyExists
	}

	if err := db.Insert(&cont); err != nil {
		return nil, err
	}

	return &cont, nil
}

// TODO: this is returning pqsql error, should return something else.
// View returns single user by ID
func (c *Contact) View(db orm.DB, id int) (*res.ContactResponse, error) {
	var contact = new(res.ContactResponse)
	sql := `SELECT "contact"."id", "contact"."name", "contact"."active", "company"."id" AS "company_id", "company"."name" AS "company_name", "contact"."profile_image", "contact"."email", 
	"contact"."birth_date", "contact"."street_name", "contact"."street_number", "cities"."id" AS "city_id", "cities"."name" AS "city_name", "states"."name" AS "state_name"
	FROM "contacts" AS "contact" LEFT JOIN "companies" AS "company" ON "company"."id" = "contact"."company_id"
	LEFT JOIN "cities" ON "cities"."id" = "contact"."city_id" 
	LEFT JOIN "states" ON "states"."id" = "cities"."state_id" 
	WHERE ("contact"."id" = ? and "contact"."deleted_at" is null)`
	_, err := db.QueryOne(contact, sql, id)
	if err != nil {
		return nil, err
	}

	return contact, nil
}

// Update updates user's contact info
func (c *Contact) Update(db orm.DB, contact *model.Contact) error {
	err := db.Update(contact)
	return err
}

// // List returns list of all users retrievable for the current user, depending on role
// func (u *User) List(db orm.DB, qp *gorsk.ListQuery, p *gorsk.Pagination) ([]gorsk.User, error) {
// 	var users []gorsk.User
// 	q := db.Model(&users).Column("user.*", "Role").Limit(p.Limit).Offset(p.Offset).Where("deleted_at is null").Order("user.id desc")
// 	if qp != nil {
// 		q.Where(qp.Query, qp.ID)
// 	}
// 	if err := q.Select(); err != nil {
// 		return nil, err
// 	}
// 	return users, nil
// }

// TODO: this should delete the contact's phone too.
// Delete sets deleted_at for a user
func (c *Contact) Delete(db orm.DB, contact *model.Contact) error {
	return db.Delete(contact)
}
