package users

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/SteveRusin/go_mentoring/middlewares"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type userRepositoryMock struct {
	mock.Mock
}

type userHandlersSuite struct {
	suite.Suite
	repositoryMock *userRepositoryMock
	handlers       userHandlers
}

func (m *userRepositoryMock) Save(user User) (*User, error) {
	args := m.Called(user)
	return args.Get(0).(*User), args.Error(1)
}

func (m *userRepositoryMock) FindByUsername(name string) (*User, error) {
	return nil, nil
}

func (m *userRepositoryMock) FindUserByCreds(creds *UserCreds) (*User, error) {
	return nil, nil
}

func (suite *userHandlersSuite) SetupTest() {
	suite.repositoryMock = new(userRepositoryMock)
	suite.handlers = userHandlers{
		repository: suite.repositoryMock,
	}
}

func (suite *userHandlersSuite) TestShouldSaveUser() {
	t := suite.T()
	url := "/user"

	payload := `
    {
      "userName": "Steve",
      "password": "qwerty"
    }
  `
	expected := &User{
		Id:       "123",
		Name:     "Steve",
		Password: "qwerty",
	}
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	suite.repositoryMock.On("Save", mock.IsType(User{})).Return(expected, nil)

	suite.handlers.User(rr, req)
	var response RegisterUserResponse

	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	if response.UserName != expected.Name || response.Id != expected.Id {
		t.Fatal("Wrong response")
	}
}

func (suite *userHandlersSuite) TestShouldReturnNotImplemented() {
	t := suite.T()
	url := "/user"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	errorResponse := suite.handlers.User(rr, req)

	if !reflect.DeepEqual(errorResponse, middlewares.NewNotImplementedError()) {
		t.Fatal("expected handler to return not implemented error")
	}
}

func TestUserHandlersSuite(t *testing.T) {
	suite.Run(t, new(userHandlersSuite))
}
