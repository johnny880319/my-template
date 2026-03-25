package mymodule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyFunction1(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "Hello, World!", MyFunction1(), "MyFunction1() should return 'Hello, World!'")
}

func TestMyFunction2(t *testing.T) {
	t.Parallel()
	assert.Equal(t, 42, MyFunction2(), "MyFunction2() should return 42")
}
