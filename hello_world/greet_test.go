package main

import "testing"

func TestGreet(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, result, expected string) {
		t.Helper()

		if result != expected {
			t.Errorf("result: %q, expected: %q", result, expected)
		}
	}

	t.Run("Greet with supplied name", func(t *testing.T) {
		result := Greet("Patrick")
		expected := "Hello, Patrick!"

		assertCorrectMessage(t, result, expected)
	})
	t.Run("Greet world when supplied empty string", func(t *testing.T) {
		result := Greet("")
		expected := "Hello, World!"

		assertCorrectMessage(t, result, expected)
	})
}
