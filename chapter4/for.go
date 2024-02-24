package main
import "fmt"
// There's no while loop in Golang for is go's while
func main(){
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i += 1
	}
	// OR
	for i:=1; i<=10; i++{
		fmt.Println(i)
	}
}

