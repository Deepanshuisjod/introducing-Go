package main

import "fmt"
import "golang-book/chapter8/mymodule"

func main(){
	xs := []float64{1,2,3,4}
	avg := mymodule.Average(xs)
	fmt.Println(avg)

}
