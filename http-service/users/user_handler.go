package users

import (
	"encoding/json"
	"net/http"

	users_rpc "github.com/SteveRusin/go_mentoring/generated"
	"github.com/SteveRusin/go_mentoring/http-service/middlewares"
)

type RegisterUserDto struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type RegisterUserResponse struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
}

type LoginUserDto = RegisterUserDto

type LoginUserResponse struct {
	Url string `json:"url"`
}

type userHandlers struct {
	userClient UserClient
}

func NewUserHandlers() *userHandlers {
	return &userHandlers{
		userClient: NewUserHTTPClient(),
	}
}

func (handler *userHandlers) User(w http.ResponseWriter, r *http.Request) *middlewares.HttpError {
	if r.Method != "POST" {
		return middlewares.NewNotImplementedError()
	}

	w.Header().Set("Content-Type", "application/json")
	registerUserDto := RegisterUserDto{}
	decodeErr := json.NewDecoder(r.Body).Decode(&registerUserDto)

	if decodeErr != nil {
		return middlewares.NewBadRequestError()
	}

	userToSave := &users_rpc.StoreUserRequest{
		Name:     registerUserDto.UserName,
		Password: registerUserDto.Password,
	}

	savedUser, err := handler.userClient.Save(userToSave)
	if err != nil {
		return middlewares.NewBadRequestError()
	}

	response := RegisterUserResponse{
		Id:       savedUser.Id,
		UserName: savedUser.Name,
	}

	json.NewEncoder(w).Encode(response)

	return nil
}

func (handler *userHandlers) UserLogin(w http.ResponseWriter, r *http.Request) *middlewares.HttpError {
	return middlewares.NewNotImplementedError()
	// if r.Method != "POST" {
	// 	return middlewares.NewNotImplementedError()
	// }
	//
	// w.Header().Set("Content-Type", "application/json")
	// loginUserDto := LoginUserDto{}
	//
	// decodeErr := json.NewDecoder(r.Body).Decode(&loginUserDto)
	//
	// if decodeErr != nil {
	// 	return middlewares.NewBadRequestError()
	// }
	//
	// _, err := handler.userClient.FindUserByCreds(&UserCreds{
	// 	Name:     loginUserDto.UserName,
	// 	Password: loginUserDto.Password,
	// })
	// if err != nil {
	// 	return middlewares.NewBadRequestError()
	// }
	//
	// response := LoginUserResponse{
	// 	Url: "ws://mock.url.io/token=.....",
	// }
	//
	// json.NewEncoder(w).Encode(response)
	//
	// return nil
}
