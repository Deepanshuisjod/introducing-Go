package main
import "fmt"

func main(){
	xs := []float64{98,93,77,82,83}
	fmt.Println(average(xs))
	fmt.Println(avg(xs))
}

// Method 1
func average(xs []float64)float64{
	total := 0.0
	for i:=0; i<len(xs); i++{
		total += xs[i]
	}
	return (total/float64(len(xs)))
}
// Method 2
func avg(xs []float64)float64{
	total := 0.0
	for _,v := range xs{
		total += v
	}
	return (total/float64(len(xs)))
}
