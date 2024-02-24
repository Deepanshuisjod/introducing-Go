package main
import "fmt"

func main(){
	var x string = "hello world"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)
	x = x + "third"
	fmt.Println(x)

	var a string = "hello"
	var b string = "hello"
	fmt.Println(a==b)

	// This this is not going to work because h is assigned int first and then it is assigned to string
	/*
	h := 12
	fmt.Println(h)
	h := "Deepanshu"
	fmt.Println(h)
	*/

	// Assigning multiple variables
	/*
	var ( l=1 o=2 p=3 )
	*/
}	

