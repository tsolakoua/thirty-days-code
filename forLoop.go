package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var n, result int
    fmt.Scan(&n)
    if 2 <= n && n <= 20 {
        for i := 1 ; i < 11; i++ {
            result = n*i
            fmt.Println(n, "x", i, "=", result)
        }   
    }
}
