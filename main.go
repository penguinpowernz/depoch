package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

// turn epochs to RFC3338 in a given text

func main() {
	// Create a scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// fmt.Println("Enter text with epoch timestamps (press Ctrl+D to finish):")

	// Regular expression to find epoch timestamps (assuming they are in seconds)
	// \b is a word boundary, \d is a digit, and {10} means exactly 10 of them.
	// So this regular expression matches 10-digit numbers which are surrounded
	// by word boundaries (i.e. not part of a larger number).
	epochRegex := regexp.MustCompile(`\b\d{10}\b`)

	// Read input line by line
	for scanner.Scan() {
		line := scanner.Text()

		// Replace function to convert epoch to RFC 3339
		output := epochRegex.ReplaceAllStringFunc(line, func(epochStr string) string {
			epoch, err := strconv.ParseInt(epochStr, 10, 64)
			if err != nil {
				return epochStr // Return the original string if parsing fails
			}
			// Convert epoch to time.Time
			t := time.Unix(epoch, 0)
			// Return the RFC 3339 formatted date
			return t.Format("2006-01-02T15:04:05")
		})

		// Output the result for the current line
		fmt.Println(output)
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}
