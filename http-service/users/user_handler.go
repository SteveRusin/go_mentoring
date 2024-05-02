package users

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	users_rpc "github.com/SteveRusin/go_mentoring/generated"
	"github.com/SteveRusin/go_mentoring/http-service/middlewares"
	"github.com/SteveRusin/go_mentoring/user-management-service/randomId"
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
	userClient    UserClient
	activeUsersDb *activeUsers
}

func NewUserHandlers() *userHandlers {
	return &userHandlers{
		userClient:    NewUserRpcClient(),
		activeUsersDb: newUsersActiveDb(),
	}
}

// https://hackernoon.com/go-the-complete-guide-to-profiling-your-code-h51r3waz
func (handler *userHandlers) ProcessImage(w http.ResponseWriter, r *http.Request) *middlewares.HttpError {
	file, err := os.Open("../upload_example.jpg")
	if err != nil {
		log.Println("Error while opening file", err)
		return middlewares.NewBadRequestError()
	}
	defer file.Close()

	hash := md5.New()

	if _, err := io.Copy(hash, file); err != nil {
		fmt.Println("Error hashing file:", err)
		return middlewares.NewBadRequestError()
	}

	hashInBytes := hash.Sum(nil)

	hashString := hex.EncodeToString(hashInBytes)

	fmt.Println("MD5 hash:", hashString)

	w.Write([]byte(hashString))

	return nil
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
	if r.Method != "POST" {
		return middlewares.NewNotImplementedError()
	}

	w.Header().Set("Content-Type", "application/json")
	loginUserDto := LoginUserDto{}

	decodeErr := json.NewDecoder(r.Body).Decode(&loginUserDto)

	if decodeErr != nil {
		return middlewares.NewBadRequestError()
	}

	res, err := handler.userClient.FindUserByCreds(&users_rpc.GetUserRequest{
		Name:     loginUserDto.UserName,
		Password: loginUserDto.Password,
	})
	if res.GetId() == "" || err != nil {
		return middlewares.NewBadRequestError()
	}

	token := randomId.New()

	handler.activeUsersDb.AddToken(loginUserDto.UserName, token)

	response := LoginUserResponse{
		Url: "ws://localhost:8080/chat?token=" + token,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		handler.activeUsersDb.RevokeToken(loginUserDto.UserName)
		log.Println("Error while encoding response", err)
		return nil
	}

	return nil
}

type ActiveUsersResponse struct {
	Users []string `json:"users"`
}

func (handler *userHandlers) GetActiveUsers(w http.ResponseWriter, r *http.Request) *middlewares.HttpError {
	if r.Method != "GET" {
		return middlewares.NewNotImplementedError()
	}

	w.Header().Set("Content-Type", "application/json")

	activeUsers := handler.activeUsersDb.GetActiveUsers()

	activeUsersResponse := ActiveUsersResponse{
		Users: activeUsers,
	}

	json.NewEncoder(w).Encode(activeUsersResponse)

	return nil
}
