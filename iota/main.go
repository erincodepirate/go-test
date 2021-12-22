package main

import "fmt"

const (
	first  = iota
	second = iota
	third  = iota + 6
	fourth
)

const (
	fifth = iota
)

func main() {
	fmt.Println(first, second)
	fmt.Println(third, fourth)
	fmt.Println(fifth)
}
