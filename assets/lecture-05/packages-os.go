package main

import (
	"log"
	"os"
)

// START OMIT

func main() {
	f, err := os.OpenFile("dump.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.WriteString("Hello from Gophers!")
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}

	data, err := os.ReadFile("dump.txt")
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(data) // Ignored error to make the example concise
	os.Remove("dump.txt")
}

// END OMIT
