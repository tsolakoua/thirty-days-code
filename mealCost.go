package main
import "fmt"

func main() {

    var mealCost float32
    var tipPercent float32
    var taxPercent float32
    
    fmt.Scan(&mealCost)
    fmt.Scan(&tipPercent)
    fmt.Scan(&taxPercent)
    
    tip := mealCost * (tipPercent/100)
    tax := mealCost * (taxPercent/100)
    totalCost := mealCost + tip + tax
    
    fmt.Print("The total meal cost is ")
    fmt.Print(int(totalCost))
    fmt.Print(" dollars.")  
}
