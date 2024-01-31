 # Project: Hasher

This project provides a simple Go package for hashing and checking passwords.

## Usage

To use this package in your Go project, follow the steps below:

1. Install the package in your project:

   ```bash
   go get github.com/SteveRusin/go_mentoring/hasher
   ```

2. Import the package in your code:

   ```go
   import "github.com/SteveRusin/go_mentoring/hasher"
   ```

3. Hash a password:

   ```go
   hashedPassword, err := hasher.HashPassword("password")
   if err != nil {
       // Handle error
   }
   // Use hashedPassword as needed
   ```

4. Check a password against its hash:

   ```go
   isValid := hasher.CheckPasswordHash("password", hashedPassword)
   // Use isValid as needed
   ```

## Example

Here's a simple example to demonstrate the usage of the `hasher` package:

```go
package main

import (
	"fmt"
	"github.com/SteveRusin/go_mentoring/hasher"
)

func main() {
	// Hash a password
	hashedPassword, err := hasher.HashPassword("mysecretpassword")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Check a password against its hash
	isValid := hasher.CheckPasswordHash("mysecretpassword", hashedPassword)

	if isValid {
		fmt.Println("Password is valid!")
	} else {
		fmt.Println("Password is invalid!")
	}
}
```

## Notes

- Ensure that you have Go installed on your machine.
- Make sure to handle errors appropriately in your code, especially when calling `HashPassword`.
- This is a simple hashing utility; for production use, consider using more advanced libraries and following best practices for password management.
