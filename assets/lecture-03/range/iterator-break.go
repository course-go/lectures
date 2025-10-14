package main

import (
	"fmt"
	"iter"
)

type Data struct {
	names []string
}

// START OMIT

func (d *Data) Iterator() iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		for i, v := range d.names {
			if !yield(i, v) {
				fmt.Println("Iterator cleanup")
				return
			}
		}
	}
}

func main() {
	ds := Data{
		names: []string{"Bob", "John", "Sarah", "William"},
	}
	for i, name := range ds.Iterator() {
		fmt.Println("Iterated value:", name, "with index:", i)
		if name == "Sarah" {
			fmt.Println("Breaking")
			break
		}
	}
}

// END OMIT
