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

type RegisterUserResponseDto struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
}

var usersRepository = NewUserRepository()

func HandleUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "404 Method not found\n")
	}

	w.Header().Set("Content-Type", "application/json")
	registerUserDto := RegisterUserDto{}
	json.NewDecoder(r.Body).Decode(&registerUserDto)

  userToSave := User{
    Id: randomId.New(),
    Name: registerUserDto.UserName,
    Password: registerUserDto.Password,
  }

  savedUser, err := usersRepository.Save(userToSave)

  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("Bad request\n"))

    return
  }

	response := RegisterUserResponseDto{
		Id:       savedUser.Id,
		UserName: savedUser.Name,
	}

  json.NewEncoder(w).Encode(response)
}
