package hasher

/*
  Hasesh password
  Usage:
    import "github.com/SteveRusin/go_mentoring/hasher"
    hasher.HashPassword("password")
*/
func HashPassword(password string) (string, error) {
    return "hashed password", nil
}

/*
  Hasesh password
  Usage:
    import "github.com/SteveRusin/go_mentoring/hasher"
    hasher.HashPassword("password")
*/
func CheckPasswordHash(password, hash string) bool {
    return true
}
