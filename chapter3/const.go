package main
import "fmt"

func main(){
	const x string = "Deepanshu"
	x = "Himanshu" // Assigning a constant will give error
	fmt.Print(x)
}
