package pznunittest

import (
	"fmt"
	"testing"
)

// setting up before and after all unit tests executed
// will be efected only same package
func TestMain(m *testing.M) {
	fmt.Println("Before executing")

	m.Run()

	fmt.Println("after executing")
}
