// Go has a special statement called defer that schedules a function call to be run afterthe function complete

package main
import "fmt"

func first(){
	fmt.Println("1st")
}

func second(){
	fmt.Println("2nd")
}
func main(){
	defer second()
	first()
}
// defer is often used when resources need to be freed in some way
