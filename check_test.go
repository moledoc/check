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
