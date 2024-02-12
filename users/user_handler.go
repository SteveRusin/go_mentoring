package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SteveRusin/go_mentoring/randomId"
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

var usersRepository = NewUserRepository()

func HandleUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	registerUserDto := RegisterUserDto{}
	decodeErr := json.NewDecoder(r.Body).Decode(&registerUserDto)

	if decodeErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request\n"))

		return
	}

	userToSave := User{
		Id:       randomId.New(),
		Name:     registerUserDto.UserName,
		Password: registerUserDto.Password,
	}

	savedUser, err := usersRepository.Save(userToSave)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request\n"))

		return
	}

	response := RegisterUserResponse{
		Id:       savedUser.Id,
		UserName: savedUser.Name,
	}

	json.NewEncoder(w).Encode(response)
}

func HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "404 Method not found\n")
	}

	w.Header().Set("Content-Type", "application/json")
	loginUserDto := LoginUserDto{}

	decodeErr := json.NewDecoder(r.Body).Decode(&loginUserDto)

	if decodeErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid username/password\n"))

		return
	}

	_, err := usersRepository.FindUserByCreds(&UserCreds{
		Name:     loginUserDto.UserName,
		Password: loginUserDto.Password,
	})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid username/password\n"))

		return
	}

	response := LoginUserResponse{
		Url: "ws://mock.url.io/token=.....",
	}

	json.NewEncoder(w).Encode(response)
}
