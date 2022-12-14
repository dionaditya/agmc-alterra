package tests

import (
	"agmc/day2/database/config"
	"agmc/day2/internal/domains"
	"agmc/day2/internal/factory"
	"agmc/day2/internal/routes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/steinfletcher/apitest"
)

var mockPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----
`

var mockPublicKey = `-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAtfSIjnrkQxg6p8SWtzdk
IcNn/lpjtlH34Eh1nmq3pmphsOlGwq+htrlqW/+yir1NXffh6kj4lpmgazvU2h/M
DUclX4dGP4aUYTF2tNlNP9WU/viDhbhOTwtq/NItnOET1QAO/zTzrMQJ39lkD9Oq
xsMgKRVkAuZvOZGnmyR0M5TywFXNFx4NB8+UmK8EEewqei+yLrsZcTaGH79ZEpaK
l6zdb0MRDXoVtSpuQ6QTnyxNwUxONw9g53j6qLqlWGwXinH+HKgGdT1swVeOS8YV
CaLhX+snmk+t9B/tYpbz75uxB1iqrIyPne1pCzRc8rIphwDfsw3wNEOkRr/+ywTW
xwxjPegs88yondZjzvKlvPDkmrUP42kz+wxe7ug0ErrcaVcu/ip48C8xNIPThzAM
Qjcv7jiiSjBVhpQW0+m590iIFTvkENLITnHHdf/6WaSNENR9TzSuh5pvCsFvE5Og
JgfYwb9q4IMir3gkKU4HQ1aYrBEFuN6csn3tcaomc/Pjt4sox5lspLbSQf0Z8eFT
gWiizSZiTsHgi8Qlp3OHe8iQscqqRbP5t27TOhcMLSjIx0Y/WJ58aXh5D3m3p/OX
8TBPQ2FgaizdpUvBZvepXsp2AS9DjB1O30RK9/6/JD/Jx+m2dfxTmacHkr4jkoc4
5YNDGEWso9GjDc9r+K6gl1ECAwEAAQ==
-----END PUBLIC KEY-----
`

func mockFactory() *factory.Factory {

	DB := config.NewDatabaseConfig(config.DatabaseConfig{
		Username:     "webuser",
		Password:     "webpass",
		Port:         "3306",
		DatabaseName: "test_db",
	})

	f := factory.NewFactory(DB)

	return f

}

func TestFailedAddNewUser(t *testing.T) {

	f := mockFactory()

	r := routes.NewRouting(f).GetRouting([]byte(mockPrivateKey), []byte(mockPublicKey))

	ts := httptest.NewServer(r)

	defer ts.Close()
	apitest.New().
		Handler(r).
		Post("/v1/users").
		Expect(t).
		Status(http.StatusBadRequest).
		End()

}

var uniqueName = uuid.NewV4()

type UserPayload struct {
	Status  string
	Message string
	Data    domains.User
}

var userPayload UserPayload

func TestAddNewUser(t *testing.T) {
	f := mockFactory()

	r := routes.NewRouting(f).GetRouting([]byte(mockPrivateKey), []byte(mockPublicKey))

	ts := httptest.NewServer(r)

	defer ts.Close()
	userData := apitest.New().
		Handler(r).
		Post("/v1/users").
		JSON(fmt.Sprintf(`{
			"Name": "%s",
			"Email": "%s@gmail.com",
			"Password": "*******"
		}`, uniqueName, uniqueName)).
		Expect(t).
		Status(http.StatusOK).
		End()

	userData.JSON(&userPayload)

}

func TestUserAlreadyExist(t *testing.T) {
	f := mockFactory()

	r := routes.NewRouting(f).GetRouting([]byte(mockPrivateKey), []byte(mockPublicKey))

	ts := httptest.NewServer(r)

	defer ts.Close()
	apitest.New().
		Handler(r).
		Post("/v1/users").
		JSON(fmt.Sprintf(`{
			"Name": "%s",
			"Email": "%s@gmail.com",
			"Password": "********"
		}`, uniqueName, uniqueName)).
		Expect(t).
		Status(http.StatusInternalServerError).
		End()

}

type Token struct {
	Token string
}

var result Token

func TestSignIn(t *testing.T) {
	f := mockFactory()

	r := routes.NewRouting(f).GetRouting([]byte(mockPrivateKey), []byte(mockPublicKey))

	ts := httptest.NewServer(r)

	defer ts.Close()
	resp := apitest.New().
		Handler(r).
		Post("/v1/signin").
		FormData("email", fmt.Sprintf("%s@gmail.com", uniqueName)).
		FormData("password", "*******").
		Expect(t).
		Status(http.StatusOK).
		End()

	resp.JSON(&result)

}

func TestErrorSignIn(t *testing.T) {
	f := mockFactory()

	r := routes.NewRouting(f).GetRouting([]byte(mockPrivateKey), []byte(mockPublicKey))

	ts := httptest.NewServer(r)

	defer ts.Close()
	apitest.New().
		Handler(r).
		Post("/v1/signin").
		FormData("email", fmt.Sprintf("%s@gmail.com", uniqueName)).
		FormData("password", "***").
		Expect(t).
		Status(http.StatusUnauthorized).
		End()

}

func TestGetUsers(t *testing.T) {
	f := mockFactory()

	r := routes.NewRouting(f).GetRouting([]byte(mockPrivateKey), []byte(mockPublicKey))

	ts := httptest.NewServer(r)

	defer ts.Close()
	apitest.New().
		Handler(r).
		Get("/v1/users").
		Header("Authorization", fmt.Sprintf("Bearer %s", result.Token)).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestFailedUnAuthGetUsers(t *testing.T) {
	f := mockFactory()

	r := routes.NewRouting(f).GetRouting([]byte(mockPrivateKey), []byte(mockPublicKey))

	ts := httptest.NewServer(r)

	defer ts.Close()
	apitest.New().
		Handler(r).
		Get("/v1/users").
		Header("Authorization", "Bearer value").
		Expect(t).
		Status(http.StatusUnauthorized).
		End()
}

func TestGetUserById(t *testing.T) {
	f := mockFactory()

	r := routes.NewRouting(f).GetRouting([]byte(mockPrivateKey), []byte(mockPublicKey))

	ts := httptest.NewServer(r)

	defer ts.Close()
	apitest.New().
		Handler(r).
		Get("/v1/users/"+userPayload.Data.ID.String()).
		Header("Authorization", fmt.Sprintf("Bearer %s", result.Token)).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestFailedNoRecordFound(t *testing.T) {
	f := mockFactory()

	r := routes.NewRouting(f).GetRouting([]byte(mockPrivateKey), []byte(mockPublicKey))

	ts := httptest.NewServer(r)

	defer ts.Close()
	apitest.New().
		Handler(r).
		Get("/v1/users/random-id").
		Header("Authorization", fmt.Sprintf("Bearer %s", result.Token)).
		Expect(t).
		Status(http.StatusNotFound).
		End()
}

func TestUpdateUser(t *testing.T) {
	f := mockFactory()

	r := routes.NewRouting(f).GetRouting([]byte(mockPrivateKey), []byte(mockPublicKey))

	ts := httptest.NewServer(r)

	defer ts.Close()
	apitest.New().
		Handler(r).
		Put("/v1/users/"+userPayload.Data.ID.String()).
		JSON(fmt.Sprintf(`{
			"Name": "%s updated",
			"Email": "%s@gmail.com",
			"Password": "****"
		}`, uniqueName, uniqueName)).
		Header("Authorization", fmt.Sprintf("Bearer %s", result.Token)).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestFailedUpdateUser(t *testing.T) {
	f := mockFactory()

	r := routes.NewRouting(f).GetRouting([]byte(mockPrivateKey), []byte(mockPublicKey))

	ts := httptest.NewServer(r)

	defer ts.Close()
	apitest.New().
		Handler(r).
		Put("/v1/users/"+userPayload.Data.ID.String()).
		JSON(fmt.Sprintf(`{
			"Unexist": "%s",
			"Email": "%s@gmail.com",
			"Password": "****"
		}`, uniqueName, uniqueName)).
		Header("Authorization", fmt.Sprintf("Bearer %s", result.Token)).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestFailedUpdateUserNoFoundUser(t *testing.T) {
	f := mockFactory()

	r := routes.NewRouting(f).GetRouting([]byte(mockPrivateKey), []byte(mockPublicKey))

	ts := httptest.NewServer(r)

	defer ts.Close()
	apitest.New().
		Handler(r).
		Put("/v1/users/lorem").
		JSON(fmt.Sprintf(`{
			"Name": "%s updated",
			"Email": "%s@gmail.com",
			"Password": "****"
		}`, uniqueName, uniqueName)).
		Header("Authorization", fmt.Sprintf("Bearer %s", result.Token)).
		Expect(t).
		Status(http.StatusInternalServerError).
		End()
}

func TestDeleteUser(t *testing.T) {
	f := mockFactory()

	r := routes.NewRouting(f).GetRouting([]byte(mockPrivateKey), []byte(mockPublicKey))

	ts := httptest.NewServer(r)

	defer ts.Close()
	apitest.New().
		Handler(r).
		Delete("/v1/users/"+userPayload.Data.ID.String()).
		Header("Authorization", fmt.Sprintf("Bearer %s", result.Token)).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestFailedDeleteUserNoUserFound(t *testing.T) {
	f := mockFactory()

	r := routes.NewRouting(f).GetRouting([]byte(mockPrivateKey), []byte(mockPublicKey))

	ts := httptest.NewServer(r)

	defer ts.Close()
	apitest.New().
		Handler(r).
		Delete("/v1/users/lorem").
		Header("Authorization", fmt.Sprintf("Bearer %s", result.Token)).
		Expect(t).
		Status(http.StatusInternalServerError).
		End()
}
