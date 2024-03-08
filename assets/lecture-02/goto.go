package main

import "fmt"

func main() {
	goto end
	fmt.Println("Why would someone do this?")
end:
	fmt.Println("I'm out!")
}
