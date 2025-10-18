package main

// START OMIT

func main() {
	increment(nil)
}

func increment(i *int) {
	*i++
}

// END OMIT
