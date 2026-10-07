package main

import (
	"fmt"
	"time"
)

var scrapedData = []string{
	"Go",
	"Is",
	"Awesome",
	"And",
	"I",
	"Can't",
	"Get",
	"Enough",
	"Of",
	"It",
}

// START OMIT

func scraper(data chan<- string) {
	// Scrape some data
	for _, sd := range scrapedData {
		fmt.Println("Scraping...")
		time.Sleep(80 * time.Millisecond)
		data <- sd // pass for processing
	}
}

func processor(in <-chan string, out chan<- string) {
	for {
		data := <-in
		fmt.Println("Processing...")
		time.Sleep(200 * time.Millisecond)
		out <- data // pass for exporting
	}
}

// MIDDLE OMIT

func exporter(data <-chan string) {
	for {
		value := <-data
		fmt.Println("Exporting...")
		time.Sleep(120 * time.Millisecond)
		fmt.Printf("Exported: %s\n", value)
	}
}

func main() {
	ingress := make(chan string, 2)
	egress := make(chan string, 2)

	go scraper(ingress)
	go processor(ingress, egress)
	go exporter(egress)

	time.Sleep(3 * time.Second)
}

// END OMIT
