package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    
    var n int
    fmt.Scan(&n)
	fmt.Print(factorial(n))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n - 1)
} 
