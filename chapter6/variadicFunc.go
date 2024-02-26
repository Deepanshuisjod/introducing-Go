package main
import "fmt"

func add(args ...int)int{
	total := 0
	for _,val := range args{
		total += val
	}
	return total
}
func main(){
	xs := []int{1,2,3,4}
	fmt.Println(add(xs ...))
}
