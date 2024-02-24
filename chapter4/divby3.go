package main
import "fmt"
func main(){
	i := 1
	for i<=100{
		if i%3 == 0{
			fmt.Print(i)
			fmt.Print(",")
		}
		i += 1
	}
}
