// check is a module that contains functions to handle errors or assert values
package check

import (
	"bufio"
	"fmt"
	_ "log"
)

// Err is a function that checks if given error is nil.
// If it is not nil, then exit with log.Fatal.
func Err(err error) {
	if err != nil {
		// log.Fatal(err)
		panic(err)
	}
}

// Scanner is a function that checks if *bufio.Scanner.Err() is nil.
// If it is not nil, then exit with log.Fatal.
func Scanner(scanner *bufio.Scanner) {
	if err := scanner.Err(); err != nil {
		// log.Fatal(err)
		panic(err)
	}
}

// Tint is a test type for int asserts.
type Tint int

// Assert is a method to check if `result` value equals to the `expected` one.
// The method panics, when the `result` does not equal to `expected` value.
func (result Tint) Assert(expected Tint) {
	if result != expected {
		panic(fmt.Sprintf("Expected %v, got %v\n", expected, result))
	}
}

// Tfloat64 is a test type for float64 asserts.
type Tfloat64 float64

// Assert is a method to check if `result` value equals to the `expected` one.
// The method panics, when the `result` does not equal to `expected` value.
func (result Tfloat64) Assert(expected Tfloat64) {
	if result != expected {
		panic(fmt.Sprintf("Expected %v, got %v\n", expected, result))
	}
}

// Tstring is a test type for string asserts.
type Tstring string

// Assert is a method to check if `result` value equals to the `expected` one.
// The method panics, when the `result` does not equal to `expected` value.
func (result Tstring) Assert(expected Tstring) {
	if result != expected {
		panic(fmt.Sprintf("Expected '%v', got '%v'\n", expected, result))
	}
}

// TintL is a test type for []int asserts.
type TintL []int

// Assert is a method to check if `result` value equals to the `expected` one.
// The method panics, when the `result` does not equal to `expected` value.
func (result TintL) Assert(expected TintL) {
	for i, elem := range result {
		if elem != expected[i] {
			panic(fmt.Sprintf("Expected %v, got %v\n", expected, result))
		}
	}
}

// Tfloat64L is a test type for []float64 asserts.
type Tfloat64L []float64

// Assert is a method to check if `result` value equals to the `expected` one.
// The method panics, when the `result` does not equal to `expected` value.
func (result Tfloat64L) Assert(expected Tfloat64L) {
	for i, elem := range result {
		if elem != expected[i] {
			panic(fmt.Sprintf("Expected %v, got %v\n", expected, result))
		}
	}
}

// TstringL is a test type for []string asserts.
type TstringL []string

// Assert is a method to check if `result` value equals to the `expected` one.
// The method panics, when the `result` does not equal to `expected` value.
func (result TstringL) Assert(expected TstringL) {
	for i, elem := range result {
		if elem != expected[i] {
			panic(fmt.Sprintf("Expected %v, got %v\n", expected, result))
		}
	}
}
