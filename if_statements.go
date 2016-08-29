package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	if 1 <= n && n <= 100 {

		if n%2 == 1 {
			fmt.Printf("Weird")

		} else if n%2 == 0 && 2 <= n && n <= 5 {
			fmt.Printf("Not Weird")
		} else if n%2 == 0 && 6 <= n && n <= 20 {
			fmt.Printf("Weird")
		} else if n > 20 {
			fmt.Printf("Not Weird")
		}

	}

}
