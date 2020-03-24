package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/goCrudChallenge/pkg/utl/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=DBName password=DBPassword") // sslmode=disable
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	//------- Insert seed values for cities, states, phone types, companies and contacts -------//

	// This creates the schema.
	db.AutoMigrate(&model.City{}, &model.Company{}, &model.Contact{}, &model.Phone{}, &model.PhoneType{}, &model.State{})

	streetName := "Street Name"
	// States
	i := 0
	for i < 5 {
		i++
		state := model.State{Name: "State " + strconv.Itoa(i)}
		db.Create(&state)
	}

	// Cities
	i = 0
	j := 1
	for i < 10 {
		i++
		city := model.City{Name: "City " + strconv.Itoa(i), StateID: j}
		// To give two cities for each state
		if i%2 == 0 {
			j++
		}
		db.Create(&city)
	}

	// Companies
	i = 0
	j = 1

	for i < 10 {
		i++
		company := model.Company{Name: "Company " + strconv.Itoa(i), CityID: i, StreetName: streetName, StreetNumber: rand.Intn(5000)}
		db.Create(&company)
	}

	// Phone types
	type1 := model.PhoneType{Name: "work"}
	type2 := model.PhoneType{Name: "personal"}
	db.Create(&type1)
	db.Create(&type2)

	// Contacts
	i = 0
	j = 1
	for i < 20 {
		i++
		contact := model.Contact{
			Name:         "Contact " + strconv.Itoa(i),
			CityID:       j,
			CompanyID:    j,
			StreetName:   streetName,
			StreetNumber: rand.Intn(5000),
			BirthDate:    randomBirthDate(),
			Email:        "contactEmail" + strconv.Itoa(i) + "@mail.com",
			ProfileImage: "ImageUrl",
		}
		// To give two contacts for each city and company
		if i%2 == 0 {
			j++
		}
		db.Create(&contact)
	}

	// Phones
	i = 0
	j = 1
	for i < 20 {
		phone := model.Phone{Prefix: 549, Number: randomPhoneNumber(), PhoneTypeID: 1, ContactID: uint(j)}
		i++
		j++

		db.Create(&phone)
	}
	i = 0
	j = 1
	for i < 20 {
		phone := model.Phone{Prefix: 549, Number: randomPhoneNumber(), PhoneTypeID: 2, ContactID: uint(j)}
		i++
		j++

		db.Create(&phone)
	}
}

func randomPhoneNumber() int {
	min := 11111111
	max := 99999999
	return rand.Intn(max-min+1) + min
}

func randomBirthDate() time.Time {
	min := 1960
	max := time.Now().Year()
	year := strconv.Itoa(rand.Intn(max-min+1) + min)
	minDay := 1
	maxDay := 28
	day := strconv.Itoa(rand.Intn(maxDay-minDay+1) + minDay)
	minMonth := 1
	maxMonth := 12
	month := strconv.Itoa(rand.Intn(maxMonth-minMonth+1) + minMonth)

	if len(day) < 2 {
		day = "0" + day
	}
	if len(month) < 2 {
		month = "0" + month
	}
	birthD, _ := time.Parse(time.RFC3339, year+"-"+month+"-"+day+"T01:00:00Z")

	return birthD
}
