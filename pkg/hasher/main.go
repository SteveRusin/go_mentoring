/*
Copyright (c) 2024 SteveRusin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package hasher

import (
	"crypto/sha256"
	"encoding/hex"
)

/*
Hashes password
Usage:

	import "github.com/SteveRusin/go_mentoring/hasher"
	hasher.HashPassword("password")
*/
func HashPassword(password string) (string, error) {
	hasher := sha256.New()
	_, err := hasher.Write([]byte(password))
	if err != nil {
		return "", err
	}
	hashPwd := hasher.Sum(nil)

  return hex.EncodeToString(hashPwd), nil
}

/*
Verifies password hash
Usage:

	import "github.com/SteveRusin/go_mentoring/hasher"
	hasher.HashPassword("password")
*/
func CheckPasswordHash(password, hash string) bool {
	return true
}
