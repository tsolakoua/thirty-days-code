package main

import "fmt"
import "bufio"
import "os"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var bookSize int
	var key string
	var value uint64
	fmt.Scan(&bookSize)

	contactBook := make(map[string]uint64)

	for i := 0; i < bookSize; i++ {

		fmt.Scan(&key)
		fmt.Scan(&value)
		contactBook[key] = value
	}


	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

    query := scanner.Text()

    if _, found := contactBook[query]; found {
			fmt.Print(query,"=",contactBook[query], "\n")
		} else {
			fmt.Println("Not found")
		}
}
	

}
