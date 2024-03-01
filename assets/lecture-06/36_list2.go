func (list *List[Value]) Print() {
	item := list.Head
	for item != nil {
		fmt.Println(item.Value)
		item = item.Next
	}
}

func main() {
	list1 := NewList[int]()
	list1.Insert(1)
	list1.Insert(2)
	list1.Insert(3)
	list1.Insert(4)
	list1.Print()

	fmt.Println()
	list2 := NewList[string]()
	list2.Insert("first")
	list2.Insert("second")
	list2.Insert("third")
	list2.Insert("fourth")
	list2.Print()
}
