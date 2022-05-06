package main

import "fmt"

const englishHelloPrefix = "Hello, "
const french = "Francais"
const frenchHelloPrefix = "Bonjour, "
const spanish = "Espanol"
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := englishHelloPrefix
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}
	return fmt.Sprintf("%s%s!", prefix, name)
}

func main() {
	fmt.Println(Hello("world", ""))
}
