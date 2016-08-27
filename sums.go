package main

import (

"fmt"
"os"
"bufio"

)

func main(){

	var i uint32 = 4;
	var d float32 = 4.0
	var s string = "Hackerrank"
	scanner := bufio.NewReader(os.Stdin)

 // Declare second integer, double, and String variables.
    var i2 uint32
    var d2 float32
    var s2 string

    // Read and save an integer, double, and String to your variables.  
    fmt.Scanf("%d", &i2)
	fmt.Scanf("%f", &d2)
    //_, err := fmt.Scanf("%f", &d2)
  //  if err != nil {err=nil }

    s2, _ = scanner.ReadString('\n')
    
    // Print the sum of both integer variables on a new line.
    fmt.Println(i+i2)
   
    // Print the sum of the double variables on a new line.

    fmt.Printf("%.1f",d+d2)
    fmt.Println()
    // Concatenate and print the String variables on a new line
    fmt.Println(s+s2)
    // The 's' variable above should be printed first.

}