package main

import "fmt"

// START OMIT

type DataStructure struct {
	items []int
}

func (ds *DataStructure) Iterator() func(yield func(int, int) bool) {
	return func(yield func(int, int) bool) {
		for i, v := range ds.items {
			if !yield(i, v) {
				return
			}
		}
	}
}

func main() {
	ds := DataStructure{
		items: []int{0, 1, 2, 3},
	}

	for i, v := range ds.Iterator() {
		fmt.Println("Value:", v, "with index:", i)
	}
}

// END OMIT
