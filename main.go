package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const arabic = "Arabic"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const arabicHelloPrefix = "Salam, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	var prefix string

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

	return prefix + name
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
