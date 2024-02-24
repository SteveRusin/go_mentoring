package user_test

import (
	"io"
	"net/http"
	"testing"

	"github.com/SteveRusin/go_mentoring/users"
	"github.com/joho/godotenv"
)

func cleanDb() {
	err := godotenv.Load(".env.integration")
	if err != nil {
		panic("Error loading .env file")
	}

	db := users.UserDBConnect()

	tx := db.Exec("TRUNCATE TABLE users;")

	if tx.Error != nil {
		panic("Error during db clean up")
	}
}

type userTests struct {
	name               string
	method             string
	body               io.Reader
	expectedStatusCode int
}

func TestUserEndpoint(t *testing.T) {
	cleanDb()
	userUrl := "http://localhost:8080/user"

	testCases := []userTests{
		{
			name:               "Should throw method now found",
			method:             http.MethodGet,
			body:               nil,
			expectedStatusCode: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(tc.method, userUrl, tc.body)
			if err != nil {
				panic(err)
			}

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				panic(err)
			}

			b, err := io.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}
			t.Log(string(b))
      if res.StatusCode != tc.expectedStatusCode {
        t.Fatal("Invalid status code")
      }
		})
	}
}
