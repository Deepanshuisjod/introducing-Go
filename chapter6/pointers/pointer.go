package main
import "fmt"

func zero(x int){
	x = 0 
}
// The pointer data type
func one(yptr *int){
	*yptr = 1
}
func main(){
	y := 4
	x := 5
	zero(x)
	fmt.Println(x) // This will still be 5
	one(&y) // Sent the address of y
	fmt.Println(y) // This will print the value 1
	// Another way to get a pointer is to use a built-in function
	z := new(int)
	one(z)
	fmt.Println(z)
	fmt.Println(*z)
}
