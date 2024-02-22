// todo mock db to test repository
package users
//
// import (
// 	"fmt"
// 	"testing"
// )
//
// var userRepository UserRepository
//
// func setUp() {
// 	userRepository = *NewUserRepository()
// }
//
// func saveMockUserToDb() {
// 	mockUser := User{
// 		Id:       "1",
// 		Name:     "Steve",
// 		Password: "123",
// 	}
//
// 	userRepository.Save(mockUser)
// }
//
// func TestShouldSaveUser(t *testing.T) {
// 	setUp()
//   saveMockUserToDb()
// 	savedUser := userRepository.db["1"]
//
// 	if savedUser.Id != "1" {
// 		t.Fatal("User not saved in db")
// 	}
// }
//
// func TestShouldReturnErrorIfUserNameIsEmpty(t *testing.T) {
// 	_, err := userRepository.FindByUsername("")
//
// 	if fmt.Sprint(err) != "User name cannot be empty" {
// 		t.Fatal("Should return an error when user name is empty string")
// 	}
// }
//
// func TestShouldFindUserByName(t *testing.T) {
//   setUp()
//   saveMockUserToDb()
//
//   user, _ := userRepository.FindByUsername("Steve")
//
//   if (*user).Id != "1" {
//     t.Fatal("User should be found by name")
//   }
// }
//
// // func TestShouldReturnNotFoundIfUserIsNotInDb(t *testing.T) {
// //   setUp()
// //
// //   user, err := userRepository.FindByUsername("Non existing")
// //
// //   if user != nil || fmt.Sprint(err) != "User not found" {
// //     t.Fatal("Should return error if user not found")
// //   }
// // }
