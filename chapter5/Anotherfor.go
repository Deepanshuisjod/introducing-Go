package main

import "fmt"

func main() {
    // Assuming x is a slice of float64 values
    x := []float64{1.2, 3.4, 5.6, 7.8}

    var total float64 = 0
    for _, value := range x {
        total += value
    }
    
    if len(x) > 0 {
        fmt.Println("Average:", total/float64(len(x)))
    } else {
        fmt.Println("Slice x is empty.")
    }
}

