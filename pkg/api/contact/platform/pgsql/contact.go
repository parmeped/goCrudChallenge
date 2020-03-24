package pgsql

import (
	"net/http"
	"strings"

	"github.com/go-pg/pg"
	"github.com/goCrudChallenge/pkg/utl/model"
	"github.com/jinzhu/gorm"
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
func (c *Contact) Create(db gorm.DB, cont model.Contact) (*model.Contact, error) {
	var contact = new(model.Contact)
	db.Where()
	err := db.Model(contact).Where("lower(name) = ? or lower(email) = ? and deleted_at is null",
		strings.ToLower(cont.Name), strings.ToLower(cont.Email)).Select()
	if (err == nil) || (err != nil && err != pg.ErrNoRows) {
		return nil, ErrAlreadyExists
	}

	if err := db.Insert(&cont); err != nil {
		return nil, err
	}

	for _, v := range cont.Phones {
		phone := model.Phone{Prefix: v.Prefix, Number: v.Number, TypeID: v.TypeID, Owner: cont.ID}
		if err := db.Insert(&phone); err != nil {
			return nil, err
		}
	}
	return &cont, nil
}

// // TODO: this is returning pqsql error, should return something else.
// // View returns single user by ID
// func (c *Contact) View(db orm.DB, id int) (*res.ContactResponse, error) {
// 	var contact = new(res.ContactResponse)
// 	sql := db.Model(&contact).Column("contact.*").Relation()

// 	//phone := model.Phone{}
// 	//q := db.Model(&contact).Column("contact.*").Join("JOIN phones p on p.owner = contact.id")
// 	// sql := db.Model(&contact).Column("contact.*, companies.name AS [company_name], cities.name AS [city_name], states.name AS [state_name]")
// 	// 		.Join("LEFT JOIN companies ON company.id = contact.company_id")
// 	// 		.Join("LEFT JOIN cities ON cities.id = contact.city_id")
// 	// 		.Join("LEFT JOIN states ON states.id = cities.state_id")
// 	// 		.Where("contact.id = ? and contact.deleted_at is null")
// 	// sql := `SELECT contact"."id", "contact"."name", "company"."id" AS "company_id", "company"."name" AS "company_name",
// 	//  "contact"."profile_image", "contact"."email", "contact"."birth_date", "contact"."street_name", "contact"."street_number",
// 	//  "cities"."id" AS "city_id", "cities"."name" AS "city_name", "states"."name" AS "state_name",
// 	//  FROM "contacts" AS "contact"
// 	//  LEFT JOIN "companies" AS "company" ON "company"."id" = "contact"."company_id"
// 	//  LEFT JOIN "cities" ON "cities"."id" = "contact"."city_id"
// 	//  LEFT JOIN "states" ON "states"."id" = "cities"."state_id"
// 	//  LEFT JOIN "phones" ON "phones"."owner" = "contact"."id"
// 	//  WHERE ("contact"."id" = ? and "contact"."deleted_at" is null)`
// 	_, err := db.QueryOne(contact, sql, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return contact, nil
// }

// // TODO: this is returning pqsql error, should return something else. Should make it return a list of the mails
// // View returns single user by ID
// func (c *Contact) ByMail(db orm.DB, mail string) (*res.ContactResponse, error) {
// 	var contact = new(res.ContactResponse)
// 	sql := `SELECT "contact"."id", "contact"."name", "contact"."active", "company"."id" AS "company_id", "company"."name" AS "company_name",
// 	 "contact"."profile_image", "contact"."email", "contact"."birth_date", "contact"."street_name", "contact"."street_number",
// 	 "cities"."id" AS "city_id", "cities"."name" AS "city_name", "states"."name" AS "state_name"
// 	 FROM "contacts" AS "contact" LEFT JOIN "companies" AS "company" ON "company"."id" = "contact"."company_id"
// 	 LEFT JOIN "cities" ON "cities"."id" = "contact"."city_id"
// 	 LEFT JOIN "states" ON "states"."id" = "cities"."state_id"
// 	 WHERE (UPPER("contact"."email") LIKE ? and "contact"."deleted_at" is null)`
// 	_, err := db.QueryOne(contact, sql, strings.ToUpper("%"+mail+"%"))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return contact, nil
// }

// // TODO: this is returning pqsql error, should return something else. Should make it return a list of the mails
// // View returns single user by ID
// func (c *Contact) ByPhone(db orm.DB, phone *req.ByPhone) (*res.ContactResponse, error) {
// 	var contact = new(res.ContactResponse)
// 	sql := `MAKE QUERY`
// 	_, err := db.QueryOne(contact, sql, "req.ByPhone")
// 	if err != nil {
// 		return nil, err
// 	}

// 	return contact, nil
// }

// // Update updates user's contact info
// func (c *Contact) Update(db orm.DB, contact *model.Contact) error {
// 	err := db.Update(contact)
// 	return err
// }

// // List returns list of all users retrievable for the current user, depending on role
// func (co *Contact) List(db orm.DB, qp *query.ListQuery, p *model.Pagination) ([]model.Contact, error) {
// 	var contacts []model.Contact
// 	q := db.Model(&contacts).Column("contact.*").Limit(p.Limit).Offset(p.Offset).Where("deleted_at is null").Order("contact.id desc")
// 	if qp != nil {
// 		q.Where(qp.Query, qp.ID)
// 	}
// 	if err := q.Select(); err != nil {
// 		return nil, err
// 	}
// 	return contacts, nil
// }

// // TODO: this should delete the contact's phone too.
// // Delete sets deleted_at for a user
// func (c *Contact) Delete(db orm.DB, contact *model.Contact) error {
// 	return db.Delete(contact)
// }
