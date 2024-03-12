package hasher_test

import (
	"fmt"
	"testing"

	"github.com/SteveRusin/go_mentoring/pkg/hasher"
)

func ExampleHasher() {
	pwd := "123456"
	fmt.Println(hasher.HashPassword(pwd))
	// Output: 8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92 <nil>
}

func BenchmarkHashPassword(b *testing.B) {
	for n := 0; n < b.N; n++ {
		hasher.HashPassword("123456")
	}
}
