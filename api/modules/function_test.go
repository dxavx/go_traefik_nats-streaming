package modules

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomString(t *testing.T) {
	var lenString = 10
	s := RandomString(lenString)
	fmt.Println(s, len(s))
	assert.Len(t, s, lenString)
}

func BenchmarkRandomString(b *testing.B) {
	RandomString(10)
}
