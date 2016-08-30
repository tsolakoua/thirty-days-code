package main

import (
    "fmt"
    "os"
    "bufio"
    
)

func main() {
    
reader := bufio.NewReader(os.Stdin)
text, _ := reader.ReadString('\n')
fmt.Println("Hello, World.")
fmt.Println(text)

}
