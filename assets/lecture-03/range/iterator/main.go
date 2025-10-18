package main

import "fmt"

// START OMIT

type Data struct {
	names []string
}

func (d *Data) Iterator() func(yield func(int, string) bool) {
	return func(yield func(int, string) bool) {
		for i, v := range d.names {
			if !yield(i, v) {
				return
			}
		}
	}
}

func main() {
	ds := Data{
		names: []string{"Bob", "John", "Sarah", "William"},
	}
	for i, v := range ds.Iterator() {
		fmt.Println("Value:", v, "with index:", i)
	}
}

// END OMIT
