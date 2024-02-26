package main
import "fmt"

func main(){
	add := func(x,y int)int{
		return x+y
	}
	fmt.Println(add(1,1))
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint){
		ret = i
		i += 2
		return
	}
}
