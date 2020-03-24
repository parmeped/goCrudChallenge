package main

import (
	"log"
	"strings"

	"github.com/goCrudChallenge/pkg/utl/model"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

func main() {
	// This is the insert that's gonna be made
	dbInsert :=
		`INSERT INTO public.cities VALUES (1, now(), now(), NULL, 'TestCity', 1);
	INSERT INTO public.states VALUES (1, now(), now(), NULL, 'TestState');
	INSERT INTO public.companies VALUES (1, now(), now(), NULL, 'TestCompany', true, 'TestStreet', 1000, 1);
	INSERT INTO public.contacts VALUES (1, now(), now(), NULL, 'TestContact', true, 1, NULL, 'test@email.com', '2020-01-01', 'TestStreet', 1000, 1);`

	var psn = `postgres://postgres:mpc3000@localhost:5432/CrudTest?sslmode=disable`
	queries := strings.Split(dbInsert, ";")

	u, err := pg.ParseURL(psn)
	checkErr(err)
	db := pg.Connect(u)
	_, err = db.Exec("SELECT 1")
	checkErr(err)

	// This creates the schema.
	createSchema(db, &model.City{}, &model.Company{}, &model.Contact{}, &model.Phone{}, &model.PhoneType{}, &model.State{})

	for _, v := range queries[0 : len(queries)-1] {
		_, err := db.Exec(v)
		checkErr(err)
	}
}

// Error checking
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// This func creates all tables
func createSchema(db *pg.DB, models ...interface{}) {
	for _, model := range models {
		checkErr(db.CreateTable(model, &orm.CreateTableOptions{
			FKConstraints: true,
		}))
	}
}
