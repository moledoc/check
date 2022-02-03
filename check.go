// check is a module that contains functions to handle errors
package check

import (
	"bufio"
	"fmt"
	"log"
)

// TODO: Comment and improve the assert method and related types

type Tint int
type Tfloat float64
type Tstring string

func (result Tint) Assert(expected Tint) {
	if result != expected {
		panic(fmt.Sprintf("Expected %v, got %v\n", expected, result))
	}
}

func (result Tfloat64) Assert(expected Tfloat64) {
	if result != expected {
		panic(fmt.Sprintf("Expected %v, got %v\n", expected, result))
	}
}

func (result Tstring) Assert(expected Tstring) {
	if result != expected {
		panic(fmt.Sprintf("Expected '%v', got '%v'\n", expected, result))
	}
}

// Err is a function that checks if given error is nil.
// If it is not nil, then exit with log.Fatal.
func Err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Scanner is a function that checks if *bufio.Scanner.Err() is nil.
// If it is not nil, then exit with log.Fatal.
func Scanner(scanner *bufio.Scanner) {
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
