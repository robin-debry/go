package main

import "testing"

func TestHello(t *testing.T) {
	name := "Karine"
	got := hello(name)
	if got == "Hello Karine" {
		return
	} else {
		t.Error("Is not Hello karine")
	}
}
