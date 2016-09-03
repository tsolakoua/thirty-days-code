package main

import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var t int 
    fmt.Scan(&t)
    
    if 1 <= t && t <= 10 {
 	
        var s string

        for i := 0; i < t; i++ {

                fmt.Scan(&s)

                if 2 <= len(s) && len(s) <= 10000 {

                    var j int
                    var oddChars, evenChars string
                    for _, r := range s {
                        if j % 2 == 1 {
                            oddChars += string(r)
                        } else {
                            evenChars += string(r)
                        }

                        j++   	
                    }
                    fmt.Println(evenChars, oddChars)  
            }
        }

    }    
}
