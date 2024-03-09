package main

func panics() {
	var i *int
	*i++
}

func main() {
	panics()
}
