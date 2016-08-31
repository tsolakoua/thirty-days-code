package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n int
	fmt.Scan(&n)

	if 1 <= n && n <= 1000 {

		arrayA := make([]int, n)

		for i := 0; i < n; i++ {
			var input int
			fmt.Scan(&input)
			arrayA[i] = input
		}

		for j := n - 1; j >= 0; j-- {
			fmt.Print(arrayA[j], " ")
		}

	}
}
