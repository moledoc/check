// check is a module that contains functions to handle errors
package check

import (
	"bufio"
	"log"
)

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
