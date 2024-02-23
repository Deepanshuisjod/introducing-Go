/*
&& and
|| or
! not
*/

package main
import "fmt"

func main(){
	var num1 uint32 = 32132
	var num2 uint32 = 42452
	fmt.Println(num1 * num2) 
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
