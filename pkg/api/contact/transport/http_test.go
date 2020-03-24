package transport_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	cs "github.com/goCrudChallenge/pkg/api/contact"
	ct "github.com/goCrudChallenge/pkg/api/contact/transport"
	"github.com/goCrudChallenge/pkg/utl/model"
	"github.com/goCrudChallenge/pkg/utl/postgres"
	"github.com/goCrudChallenge/pkg/utl/server"
)

type casesStruct struct {
	name       string
	req        string
	wantStatus int
	wantResp   *model.Contact
}

var createCases = []casesStruct{
	{
		name: "Fail on name length",
		req: `
						{
							"name" : "t",
							"email": "testMail@email.com",
							"profile_image": "ImageURL",
							"birth_date": "2000-01-01T00:00:00Z",
							"company_id": 1,
							"street_name": "Name Of Street",
							"street_number": 1100,
							"city_id": 10,
							"phones": [
								{
									"prefix": 911,
									"number": 11111111,
									"type_id": 1
								},
								{
									"prefix": 911,
									"number": 99999999,
									"type_id": 2
								}]
						}
			`,
		wantStatus: http.StatusBadRequest,
	},
	{
		name: "Fail on worng mail",
		req: ` 
			{
				"name" : "t",
				"email": "testMail@email",
				"profile_image": "ImageURL",
				"birth_date": "2000-01-01T00:00:00Z",
				"company_id": 1,
				"street_name": "Name Of Street",
				"street_number": 1100,
				"city_id": 10,
				"phones": [
					{
						"prefix": 911,
						"number": 11111111,
						"type_id": 4
					},
					{
						"prefix": 911,
						"number": 99999999,
						"type_id": 2
					}]
			}
		`,
		wantStatus: http.StatusBadRequest,
	},
	{
		name: "Fail on phone type",
		req: ` 
			{
				"name" : "t",
				"email": "testMail@email.com",
				"profile_image": "ImageURL",
				"birth_date": "2000-01-01T00:00:00Z",
				"company_id": 1,
				"street_name": "Name Of Street",
				"street_number": 1100,
				"city_id": 10,
				"phones": [
					{
						"prefix": 911,
						"number": 11111111,
						"type_id": 4
					},
					{
						"prefix": 911,
						"number": 99999999,
						"type_id": 2
					}]
			}
		`,
		wantStatus: http.StatusBadRequest,
	},
	{
		name: "Create success",
		req: ` 
			{
				"name" : "testing",
				"email": "testMailOk@email.com",
				"profile_image": "ImageURL",
				"birth_date": "2000-01-01T00:00:00Z",
				"company_id": 1,
				"street_name": "Name Of Street",
				"street_number": 1100,
				"city_id": 10,
				"phones": [
					{
						"prefix": 911,
						"number": 11111111,
						"type_id": 1
					},
					{
						"prefix": 911,
						"number": 99999999,
						"type_id": 2
					}]
			}
		`,
		wantStatus: http.StatusOK,
	},
}

// TODO: [IMPORTANT IMPROVEMENT]: THIS SHOULD ALL BE MOCKED. It will fail after the first success, aside from the fact that it's using a "real" db.
func TestCreate(t *testing.T) {

	for _, tt := range createCases {

		db, err := postgres.New("host=localhost port=5432 user=postgres dbname=CrudTest password=mpc3000 sslmode=disable")
		if err != nil {
			t.Fatal(err)
		}
		e := server.New()
		v1 := e.Group("/v1")
		ct.NewHTTP(cs.Initialize(db), v1)

		path := "http://localhost:8080/v1/contacts"

		res, err := http.Post(path, "application/json", bytes.NewBufferString(tt.req))
		if err != nil {
			t.Fatal(err)
		}
		defer res.Body.Close()
		if tt.wantResp != nil {
			response := new(model.Contact)
			if err := json.NewDecoder(res.Body).Decode(response); err != nil {
				t.Fatal(err)
			}
			if response != tt.wantResp {
				t.Error(tt.name)
			}
		}
		if res.StatusCode != tt.wantStatus {
			t.Error(tt.name)
		}
	}

}
