package modules

import (
	"math/rand"
)

func RandomString(n int) string {

	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	s := make([]byte, n)

	for i := 0; i < n; i++ {
		t := rand.Intn(len(letters))
		s[i] = letters[t]
	}
	return string(s)
}
