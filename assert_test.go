package assert

import (
	"fmt"
	"testing"
)

// func ExampleEqual() {
// t := &testing.T{}
func TestEqual(t *testing.T) {
	t.Run("call t.Fatal() on error", func(t *testing.T) {
		fmt.Println("Will always print this")
		Equal(t, 10, 11, "Error message")

		fmt.Println("This print will never run")
	})

	t.Run("works for any type", func(t *testing.T) {
		Equal(t, struct{ a int }{10}, struct{ a int }{11}, "Error message")
	})
}

func TestNotEqual(t *testing.T) {
	t.Run("call t.Fatal() on error", func(t *testing.T) {
		fmt.Println("Will always print this")
		NotEqual(t, 10, 10, "Error message")

		fmt.Println("This print will never run")
	})

	t.Run("works for any type", func(t *testing.T) {
		NotEqual(t, struct{ a int }{10}, struct{ a int }{10}, "Error message")
	})
}
