package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	switch language {
	case spanish:
		return spanishHelloPrefix + name
	case french:
		return frenchHelloPrefix + name
	default:
		return englishHelloPrefix + name
	}
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
