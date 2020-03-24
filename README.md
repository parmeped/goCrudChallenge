# --  goCrudChallenge --

# Tech Stack
goLang 
gorm
postgresql

# Quick Start
To get dependencies : go get ./...

To create the db, go to goCrudChallenge/cmd/migration, then execute -- > go run main.go

To start the api, go to goCrudChallenge/cmd/api, then execute -- > go run main.go

This starts a server on localhost:8080/

# Endpoints
The following endpoints are available: 

## Create
POST localhost:8080/v1/contacts

JSON:
{
	"name" : "Testing",
	"email": "testingAnEmail@email.com",
	"profile_image": "ImageURL",
	"birth_date": "2000-01-01T00:00:00Z", 
	"company_id": 1,
	"street_name": "Street Name",
	"street_number": 1100,
	"city_id": 10,
	"phones": [
		{
			"prefix": 911,
			"number": 20201111,
			"type_id": 4
		},
		{
			"prefix": 911,
			"number": 40409090,
			"type_id": 2
		}]
}

Returns a 200 ok if saved correctly. 

Otherwise, a 500 Internal Server Error if type_id is incorrect (accepted 1, 2).
A 400 Bad Request is returned if the mail or name for the contact already exist, if the company or city don't exist, if the name doesn't comply with a min lenght, or if the email isn't valid.

## View
GET localhost:8080/v1/contacts/:id

Returns a 200 ok with the contact info if found.

Otherwise, a 500 Internal Server Error if not found and 400 Bad Request if a string is passed.

## Delete 
DELETE localhost:8080/v1/contacts/:id

Returns a 200 ok if deleted. Performs a soft delete on the contact and phone numbers. 

Otherwise, a 500 Internal Server Error if not found and 400 Bad Request if a string is passed.

## Update
PATCH localhost:8080/v1/contacts/:id 

JSON: 
{
	"name" : "Testing Update!",
	"email": "testOfAnUpdate@email.com",
	"profile_image": "TestImageUpdated",
	"birth_date": "2080-01-01T15:00:00Z",
	"company_id": 10,
	"street_name": "Awesome street name",
	"street_number": "10",
	"city_id": 1,
	"state_id": 1
}

Returns a 200 ok if updated, and the new data is shown.

Otherwise, a 500 Internal Server Error if not found, or 400 Bad Request if company or city don't exist.

# List by Location (Company or City)
GET localhost:8080/v1/contacts/listByLocation/:searchParam/:id

JSON: 
{
	"Limit": 5,
	"Page": 0
}

Returns a 200 ok with the paginated data when found. 
If a wrong parameter is passed, returns a 500.

# List by email 
GET localhost:8080/v1/contacts/listByMail/:mail

JSON: 
{
	"Limit": 5,
	"Page": 0
}

Returns a 200 ok with the paginated data when found.

# By Phone
GET localhost:8080/v1/contacts/listByPhone

JSON: 
{
	"prefix": 20,
	"number": 57663493
}

Returns a 200 ok with the data, when found. 
Returns a 500 Internal Server Error when not found.