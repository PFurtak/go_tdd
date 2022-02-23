package main

import "testing"

func TestGreet(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, result, expected string) {
		t.Helper()

		if result != expected {
			t.Errorf("result: %q, expected: %q", result, expected)
		}
	}

	t.Run("Greet with supplied name in english", func(t *testing.T) {
		result := Greet("Patrick", "english")
		expected := "Hello, Patrick!"

		assertCorrectMessage(t, result, expected)
	})

	t.Run("Greet with supplied name in Spanish", func(t *testing.T) {
		result := Greet("Patrick", "spanish")
		expected := "Hola, Patrick!"

		assertCorrectMessage(t, result, expected)
	})

	t.Run("Greet with supplied name in French", func(t *testing.T) {
		result := Greet("Patrick", "french")
		expected := "Bonjour, Patrick!"

		assertCorrectMessage(t, result, expected)
	})

	t.Run("Greet world in english when supplied empty string for name and language", func(t *testing.T) {
		result := Greet("", "")
		expected := "Hello, World!"

		assertCorrectMessage(t, result, expected)
	})
}
