package main

type List[T any] struct {
	Head *Item[T]
}

type Item[T any] struct {
	Value T
	Next  *Item[T]
}

func (list *List[T]) Insert(value T) {
	item := Item[T]{
		Value: value,
	}
	item.Next = list.Head
	list.Head = &item
}
