package main

import (
	"DataStructures_golang/List"
	"fmt"
)

func main() {
	list := List.List[int]{}
	list.Init(1)
	list.Append(2, 3, 4, 5)

	fmt.Println(list.GetByIndex(1))

	if err := list.DeleteByIndex(1); err != nil {
		fmt.Println(err)
	}

	fmt.Println(list.GetByIndex(1))
}
