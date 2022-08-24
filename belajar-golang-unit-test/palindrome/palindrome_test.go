package palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	IsPalindrome("eye")
}

func TestIsPalindrome(t *testing.T) {

	assert.True(
		t,
		IsPalindrome("eye"),
		"Must be true and palindrome",
	)

	assert.False(
		t,
		IsPalindrome("car"),
		"Must be false and not palindrome",
	)
}
