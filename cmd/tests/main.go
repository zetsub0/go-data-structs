package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/zetsub0/go-data-structs/set"
)

func main() {
	s := set.New[int]()

	go add(s)
	go contains(s)
	go print(s)
	go remove(s)

	time.Sleep(3 * time.Second)
	fmt.Println("\n\n\n\n")
	fmt.Println(s.Length())
}

func add(s *set.Set[int]) {
	for range 1000 {
		s.Add(rand.Intn(1000))
	}
}

func remove(s *set.Set[int]) {
	for range 1000 {
		s.Remove(rand.Intn(1000))
	}
}

func contains(s *set.Set[int]) {
	for range 1000 {
		s.Contains(rand.Intn(1000))
	}
}

func print(s *set.Set[int]) {
	for range 1000 {
		s.Print()
	}
}
