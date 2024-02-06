# Go Mentoring

To use the provided Go modules for managing messages and users, follow these steps:

1. **Importing Modules**: Import the modules into your Go project where you intend to use them.

```go
import (
	"github.com/SteveRusin/go_mentoring/messages"
	"github.com/SteveRusin/go_mentoring/users"
)
```

2. **Creating Repositories**: Initialize repositories for messages and users.

```go
messageRepo := messages.NewMessagRepository()
userRepo := users.NewUserRepository()
```

3. **Saving Users and Messages**:

```go
// Creating a new user
newUser := users.User{
	Id:       "1",
	Name:     "John",
	Password: "123456",
}
userRepo, err := userRepo.Save(newUser)
if err != nil {
    // handle error
}

// Saving a message for a user
newMessage := messages.Message{
	Id:   "1",
	Text: "Hello, world!",
}
err = messageRepo.Save(newUser.Id, newMessage)
if err != nil {
    // handle error
}
```

4. **Finding Users and Messages**:

```go
// Finding a user by username
foundUser, err := userRepo.FindByUsername("John")
if err != nil {
    // handle error
}
fmt.Println("Found user:", foundUser)

// Finding all messages for a user
allMessages, err := messageRepo.findAll(foundUser.Id)
if err != nil {
    // handle error
}
for _, msg := range allMessages {
    fmt.Println("Message:", msg)
}
```

5. **Running Tests**: You can run tests for both modules using the following commands:

For messages:
```sh
go test -v ./messages
```

For users:
```sh
go test -v ./users
```

Remember to replace `"github.com/SteveRusin/go_mentoring/messages"` and `"github.com/SteveRusin/go_mentoring/users"` with the actual paths where you have stored these modules.

**Note**: Ensure that you have set up your Go workspace correctly and have the necessary dependencies installed.

Feel free to customize the provided code according to your specific requirements and application structure.
