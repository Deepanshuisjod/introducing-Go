package main
import "fmt"

func main(){
	x := 12
	y := 12
	fmt.Println(sum(x,y))
}

func sum(x int, y int)int{
	return (x+y)
}
