package main

import "testing"

func TestGreet(t *testing.T) {
	t.Run("Greet with supplied name", func(t *testing.T) {
		got := Greet("Patrick")
		expected := "Hello, Patrick!"

		if got != expected {
			t.Errorf("got: %q, expected: %q", got, expected)
		}
	})
	t.Run("Greet world when supplied empty string", func(t *testing.T) {
		got := Greet("")
		expected := "Hello, World!"

		if got != expected {
			t.Errorf("got: %q, expected: %q", got, expected)
		}
	})

}
