package main

import (
	"github.com/goCrudChallenge/pkg/utl/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=CrudTest password=mpc3000 sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// This creates the schema.
	db.AutoMigrate(&model.City{}, &model.Company{}, &model.Contact{}, &model.Phone{}, &model.PhoneType{}, &model.State{})

}
