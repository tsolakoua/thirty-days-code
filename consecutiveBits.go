package main
import "fmt"
import "strconv"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var n int64
    fmt.Scan(&n)
    max:= 0
    counter := 0
    binary_number := strconv.FormatInt(n, 2) // binary_number is string

 		for _, element := range binary_number {

               if string(element) == "1" {
               	counter++
               }	else {
               		if counter > max {
               			max = counter
               		}
               		counter = 0
               	}
        }

        if counter > max {
          fmt.Println(counter)
        } else {
          fmt.Println(max)
        }             
}
