package main

import "fmt"

func main() {
	greetings := hello("Karine")
	fmt.Println(greetings)
}

func hello(name string) string {
	return "Hello " + name
}
