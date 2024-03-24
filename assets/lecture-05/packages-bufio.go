package main

import (
	"bufio"
	"fmt"
	"os"
)

// START OMIT

func main() {
	writer := bufio.NewWriter(os.Stdout) // `os.Stdout` satisfies the `io.Writer` interface
	fmt.Fprintln(writer, "Hello\nGophers!")

	text := "	~ course-go"
	written, err := writer.WriteString(text)
	if err != nil || written != len(text) {
		fmt.Fprintln(os.Stderr, "failed writing to stdout or the text is incomplete")
	}

	writer.Flush()

	scanner := bufio.NewScanner(os.Stdin) // `os.Stdin` satisfies the `io.Reader` interface
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

// END OMIT
