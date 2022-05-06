package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanish = "Espanol"
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return fmt.Sprintf("%s%s!", spanishHelloPrefix, name)
	}
	return fmt.Sprintf("%s%s!", englishHelloPrefix, name)
}

func main() {
	fmt.Println(Hello("world", ""))
}
