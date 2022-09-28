package main

import "fmt"

func main() {
	greetings := hello("Victor")
	fmt.Println(greetings)
}

func hello(name string) string {
	return "Hello " + name
}
