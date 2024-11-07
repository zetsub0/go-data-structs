package main

import "fmt"

func main() {
	mainSlice := []int{1, 2, 3, 4, 5}
	mainSlice = append(mainSlice, 6)
	fmt.Println(len(mainSlice), cap(mainSlice))
	r(mainSlice)
	r(mainSlice)
	r(mainSlice)
	r(mainSlice)
	r(mainSlice)

	fmt.Println(len(mainSlice), cap(mainSlice))
	fmt.Println(mainSlice[:cap(mainSlice)])
}

func r(s []int) {
	s = append(s, 7)
	fmt.Println(s)
}
