package check_test

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/moledoc/check"
)

func TestCheck(t *testing.T) {
	file, err := os.Open("test_scanner.csv")
	defer file.Close()
	check.Err(err)
	scanner := bufio.NewScanner(file)
	defer check.Scanner(scanner)
	for scanner.Scan() {
		kv := strings.Split(scanner.Text(), ",")
		fmt.Println(kv)
	}
}

func TestAssertI(t *testing.T) {
	defer func() {
		expected := "Expected 2, got 1\n"
		if err := recover(); err != nil && err != expected {
			t.Fatalf("Expected '%v', got '%v'\n", expected, err)
		}
	}()
	check.Tint(1).Assert(2)
}

func TestAssertF(t *testing.T) {
	defer func() {
		expected := "Expected 2, got 1\n"
		if err := recover(); err != nil && err != expected {
			t.Fatalf("Expected '%v', got '%v'\n", expected, err)
		}
	}()
	check.Tfloat64(1).Assert(2)
}

func TestAssertS(t *testing.T) {
	defer func() {
		expected := "Expected '2', got '1'\n"
		if err := recover(); err != nil && err != expected {
			t.Fatalf("Expected '%v', got '%v'\n", expected, err)
		}
	}()
	check.Tstring("1").Assert("2")
}
