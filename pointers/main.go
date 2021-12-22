package main

import "fmt"

func main() {
	var firstName *string = new(string)

	*firstName = "foo"

	fmt.Println(*firstName)

	fname := "foo"
	ptr := &fname

	fmt.Println(ptr, *ptr)

	fname = "fee"

	fmt.Println(ptr, *ptr)
}
