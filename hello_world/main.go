package main

import "fmt"

func createLocalePrefix(language string) (prefix string) {
	switch language {
	case "spanish":
		prefix = "Hola, "
	case "french":
		prefix = "Bonjour, "
	default:
		prefix = "Hello, "

	}
	return
}

func Greet(name string, language string) string {

	if name == "" {
		name = "World"
	}

	localizedPrefix := createLocalePrefix(language)

	return localizedPrefix + name + "!"
}

func main() {
	fmt.Println(Greet("Patrick", "english"))
}
