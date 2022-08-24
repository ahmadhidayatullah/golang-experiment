package pznunittest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubTest(t *testing.T) {
	t.Run("Hi", func(t *testing.T) {

		assert.Equal(t, "Hi", "Hi")
	})
}
