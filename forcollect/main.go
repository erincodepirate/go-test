package main

func main() {
	slice := []int{1, 2, 3}

	for i, v := range slice {
		println(i, v)
	}

	ports := map[string]int{"http": 80, "https": 443}
	for k, v := range ports {
		println(k, v)
	}
}
