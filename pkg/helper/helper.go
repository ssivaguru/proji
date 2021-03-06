package helper

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

// DoesPathExist checks if a given path exists in the filesystem.
func DoesPathExist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// StrToUInt converts a string into a uint.
func StrToUInt(num string) (uint, error) {
	// Parse the input
	id64, err := strconv.ParseUint(num, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(id64), nil
}

// WantTo waits for a valid user input to confirm if he wants to do whatever was asked for.
func WantTo(question string) bool {
	var input string
	for {
		fmt.Print(question + " [y/N] ")
		n, err := fmt.Scan(&input)
		if n == 1 && err == nil {
			input = strings.ToLower(input)
			if input == "n" || input == "\n" {
				return false
			} else if input == "y" {
				return true
			}
		}
	}
}

// IsInSlice returns true if a given string is found in the given slice and false if not.
func IsInSlice(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// SkipNetworkBasedTests skips network/internet dependent tests when the env variable PROJI_SKIP_NETWORK_TESTS is set to 1
func SkipNetworkBasedTests(t *testing.T) {
	env := os.Getenv("PROJI_SKIP_NETWORK_TESTS")
	if env == "1" {
		t.Skip("Skipping network based tests")
	}
}
