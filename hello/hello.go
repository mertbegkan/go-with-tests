package main

import "fmt"

const turkish = "Turkish"
const french = "French"
const englishHelloPrefix = "Hello, "
const turkishhHelloPrefix = "Merhaba, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case turkish:
		prefix = turkishhHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Mert", ""))
}
