package main

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func main() {
	u1 := User{
		ID:        1,
		FirstName: "Nadya",
		LastName:  "Krupskaya",
	}
	u2 := User{
		ID:        2,
		FirstName: "Frida",
		LastName:  "Kahlo",
	}
	if u1 == u2 {
		println("Same user")
	} else if u1.FirstName == u2.FirstName {
		println("Similar user")
	} else {
		println("Different users")
	}

}
