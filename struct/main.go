package main

import "fmt"

func main() {
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user

	u.ID = 1
	u.FirstName = "Bob"
	u.LastName = "Ross"
	fmt.Println(u)

	u2 := user{ID: 1, FirstName: "Thomas", LastName: "Sankara"}
	fmt.Println(u2)

	u3 := user{
		ID:        2,
		FirstName: "Karl",
		LastName:  "Marx",
	}
	fmt.Println(u3)
}
