package main

import (
	"fmt"
	"unsafe"
)

// START OMIT

type product struct {
	name  string
	stock int32
	price uint64
}

func main() {
	product := product{name: "Flamethrower", stock: -72, price: 999}
	fmt.Println("Size:", unsafe.Sizeof(product))
	fmt.Println("Alignment:", unsafe.Alignof(product))
	fmt.Println("Name offset:", unsafe.Offsetof(product.name))
	fmt.Println("Price offset:", unsafe.Offsetof(product.price))

	productPointer := unsafe.Pointer(&product)
	stock := (*int32)(unsafe.Add(productPointer, unsafe.Offsetof(product.stock)))
	fmt.Println("Stock", *stock)

	truth := []string{"Go", "is", "the", "best!"}
	fmt.Println("Slice size:", unsafe.Sizeof(truth))
	slicePointer := unsafe.Pointer(unsafe.SliceData(truth))
	itemSize := unsafe.Sizeof("")
	for i := range len(truth) {
		fmt.Println(*(*string)(unsafe.Add(slicePointer, uintptr(i)*itemSize)))
	}
}

// END OMIT
