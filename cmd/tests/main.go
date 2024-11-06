package main

import (
	"log"

	"github.com/zetsub0/go-data-structs/pkg/linkedList"
)

func main() {
	ll := linkedList.New[any]([]any{any([]int{1, 2, 3, 4, 5}), any([]string{"asd", "bvf"})})
	if err := ll.RemoveTail(); err != nil {
		log.Println(err)
	}
	if err := ll.RemoveTail(); err != nil {
		log.Println(err)
	}
	if err := ll.RemoveTail(); err != nil {
		log.Println(err)
	}
	if err := ll.RemoveTail(); err != nil {
		log.Println(err)
	}
	if err := ll.RemoveTail(); err != nil {
		log.Println(err)
	}
	if err := ll.RemoveTail(); err != nil {
		log.Println(err)
	}
	if err := ll.RemoveTail(); err != nil {
		log.Println(err)
	}
	if err := ll.RemoveTail(); err != nil {
		log.Println(err)
	}
	if err := ll.RemoveTail(); err != nil {
		log.Println(err)
	}

	if err := ll.RemoveHead(); err != nil {
		log.Println(err)
	}

	if err := ll.RemoveNode(12); err != nil {
		log.Println(err)
	}

	if err := ll.RemoveNode(-1); err != nil {
		log.Println(err)
	}
}
