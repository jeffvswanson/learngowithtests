package main

import "fmt"

const englishHelloPrefix = "Hello, "
const french = "Francais"
const frenchHelloPrefix = "Bonjour, "
const german = "Deutsch"
const germanHelloPrefix = "Guten Tag, "
const spanish = "Espanol"
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := greetingPrefix(language)
	return fmt.Sprintf("%s%s!", prefix, name)
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case german:
		prefix = germanHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
