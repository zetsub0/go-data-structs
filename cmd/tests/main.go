package main

import (
	"github.com/zetsub0/go-data-structs/pkg/linkedList"
)

func main() {
	ll := linkedList.New[int]()
	ll.ParseArray([]int{1, 2, 3, 4, 5})
	ll.Print()
}
