package main

import (
	"fmt"

	"github.com/erincodepirate/go-test/webservice/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Rosa",
		LastName:  "Luxemburg",
	}
	fmt.Println(u)
}
