package main

import "testing"

func TestHello(t *testing.T) {
	name := "Karine"
	got := hello(name)
	if got != "Hello "+name {
		t.Error("Is not Hello " + name)
	}
}
