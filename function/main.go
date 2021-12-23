package main

import (
	"fmt"
)

func main() {
	port := 3000
	_, err := startWebServer(port, 2)
	fmt.Println(err)
}

func startWebServer(port, retries int) (int, error) {
	fmt.Println("Starting server")
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", retries)

	//return errors.New("Something went wrong")
	return port, nil
}
