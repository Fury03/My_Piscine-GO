package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// If arguments are provided → treat them as file paths
	if len(os.Args) > 1 {
		for _, filename := range os.Args[1:] {
			data, err := os.ReadFile(filename)
			if err != nil {
				printStr("ERROR: " + err.Error() + "\n")
				os.Exit(1)
			}
			printStr(string(data))
		}
		return
	}

	// Otherwise: read from stdin
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		printStr("ERROR: " + err.Error() + "\n")
		os.Exit(1)
	}
	printStr(string(data))
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
