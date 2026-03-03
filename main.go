package main

import (
	"fmt"
	"io"
	"os"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

func main() {
	var input io.Reader = os.Stdin
	var outputFile *os.File

	args := os.Args[1:]

	if len(args) >= 1 {
		f, err := os.Open(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error opening input file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		input = f
	}

	if len(args) >= 2 {
		f, err := os.Create(args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error opening output file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		outputFile = f
	}

	htmlBytes, err := io.ReadAll(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}

	converter := md.NewConverter("", true, nil)
	markdown, err := converter.ConvertString(string(htmlBytes))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error converting HTML: %v\n", err)
		os.Exit(1)
	}

	if outputFile != nil {
		fmt.Fprintln(outputFile, markdown)
	} else {
		fmt.Println(markdown)
	}
}
