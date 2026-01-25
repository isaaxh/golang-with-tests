package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"
	arabic  = "Arabic"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	arabicHelloPrefix  = "Salam, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case arabic:
		prefix = arabicHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Tom", french))
}

// TODO how to TDD
// 1. write a test for the requirment - done
// 2. make the compiler pass - done
// 3. run the test, see that it fails and check the error message is meaningful
// 4. write enough code to make the test pass again
// 5. refactor
