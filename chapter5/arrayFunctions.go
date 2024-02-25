package main
import "fmt"

func main(){
	fmt.Println("Enter the elements in the array ")
	// Append function
	arr0 := []int{1,2,3}
	arr1 := append(arr0,100)
	fmt.Println("New array :",arr1)

	// Copy function
	x := []int{1,2,3}
	y := make([]int,5)
	z := []int{4,5,6,7}
	copy(y,x) // This will copy the contents of x to y
	copy(z,x) // This will copy the contents of z to y
	fmt.Println(x,y)
	fmt.Println(z)
}

