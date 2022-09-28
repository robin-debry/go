package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Karine", func(t *testing.T) {
		name := "Karine"

		got := hello(name, "en")
		if got != "Hello "+name {
			t.Error("Is not Hello " + name)
		}
	})
	t.Run("Salah", func(t *testing.T) {
		name := "Salah"

		got := hello(name, "en")
		if got != "Hello "+name {
			t.Error("Is not Hello " + name)
		}
	})

	t.Run("Salah", func(t *testing.T) {
		name := "Salah"

		got := hello(name, "fr")
		if got != "Bonjour "+name {
			t.Error("Is not bonjour " + name)
		}
	})
	t.Run("Salah", func(t *testing.T) {
		name := "Salah"

		got := hello(name, "es")
		if got != "Hola "+name {
			t.Error("Is not Hola " + name)
		}
	})

}
