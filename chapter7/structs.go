package main
import (
	"fmt"
	"math"
)

type Circle struct{
	x, y, r float64
}

func circleArea(c *Circle) float64{ // to change the value of x we would simply apply an astrick(*)
	c.x = 11
	return math.Pi * c.r*c.r
}

func main(){
	var c Circle // This is the instance of our struct Circle
	// Or we can use the new function
	// c := new(Circl) <- This returns a pointer to the struct (*Circle)
	//c := Circle{x: 0, y: 0, r: 5}
	c.x = 10
	c.y = 5
	c.r = 2.5
	fmt.Println(circleArea(c))
	fmt.Println(c.x)
}
