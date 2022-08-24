package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Ahmad")

	// if result != "Hello1 Ahmad" {
	// 	// panic("Result is not Ahmad")
	// 	// t.Fail()
	// 	t.Error("Result is not Ahmad")
	// }

	assert.Equal(t, "Hello Ahmad", result, "They must be same")
	assert.True(t, false, "It must be true")
	//assertion library (Testify)
}
