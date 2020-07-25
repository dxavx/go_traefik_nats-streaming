package modules

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomeString(t *testing.T) {
	var lenString = 10
	RandomeString(lenString)
	assert.Len(t, RandomeString(lenString), lenString)
}

func BenchmarkRandomeString(b *testing.B) {
	var lenString = 10
	RandomeString(lenString)
	assert.Len(b, RandomeString(lenString), lenString)
}
