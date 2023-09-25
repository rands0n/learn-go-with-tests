package main

import "fmt"

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	prefix := "Hello, "

	switch lang {
	case "Spanish":
		prefix = "Hola, "
	case "French":
		prefix = "Bonjour, "
	}

	return prefix + name + "!"
}

func main() {
	fmt.Println(Hello("Rands", ""))
}
