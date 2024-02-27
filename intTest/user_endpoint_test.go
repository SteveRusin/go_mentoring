package user_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/SteveRusin/go_mentoring/users"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func cleanDb(t *testing.T) {
	err := godotenv.Load(".env.integration")
	if err != nil {
		t.Fatal("Error loading .env file")
	}

	db := users.UserDBConnect()

	tx := db.Exec("TRUNCATE TABLE users;")

	if tx.Error != nil {
		t.Fatal("Error during db clean up")
	}
}

type userTests struct {
	name               string
	method             string
	body               io.Reader
	expectedStatusCode int
	expectedBody       string
}

func TestUserEndpoint(t *testing.T) {
	userUrl := "http://localhost:8080/user"

	testCases := []userTests{
		{
			name:               "Should throw method now found",
			method:             http.MethodGet,
			body:               nil,
			expectedStatusCode: http.StatusNotImplemented,
		},
		{
			name:   "Should register user",
			method: http.MethodPost,
			body: strings.NewReader(`
        {
          "userName": "Test",
          "password": "123"
        }
      `),
			expectedStatusCode: http.StatusOK,
      expectedBody: "\"userName\":\"Test\"",
		},
	}

	for _, tc := range testCases {
		cleanDb(t)
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(tc.method, userUrl, tc.body)
			if err != nil {
				t.Fatal(err)
			}

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatal(err)
			}

			b, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}

			if res.StatusCode != tc.expectedStatusCode {
				t.Fatal("Invalid status code")
			}

			if tc.expectedBody != "" {
				assert.Contains(t, string(b), tc.expectedBody)
			}
		})
	}
}
