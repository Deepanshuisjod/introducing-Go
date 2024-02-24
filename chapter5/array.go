package main
import "fmt"

func main(){
	// Array can also be created as :
	a := [5]float64{ 98, 93, 77, 82, 83 }
	fmt.Println(a)
	var x[5] int
	x[4] = 100
	fmt.Println(x) // [0,0,0,0,100] <- answer
	fun()
}

// This function will print the average 
func fun(){
	var y[5] float64
	y[0] = 98
	y[1] = 93
	y[2] = 77
 	y[3] = 82
	y[4] = 83
	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += y[i]
	}
	fmt.Println(total / 5)
}
