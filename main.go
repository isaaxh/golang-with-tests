package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello(""))
}

// TODO how to TDD
// 1. write a test for the requirment
// 2. make the compiler pass
// 3. run the test, see that it fails and check the error message is meaningful
// 4. write enough code to make the test pass again
// 5. refactor
