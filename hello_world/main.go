package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Greet(name string) string {

	if name == "" {
		name = "World"
	}

	return englishHelloPrefix + name + "!"
}

func main() {
	fmt.Println(Greet("Patrick"))
}
