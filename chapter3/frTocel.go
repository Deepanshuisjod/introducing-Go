package main
import "fmt"

func main(){
	fmt.Print("Enter the temperature in Fahrenheit : ")
	var F float64
	fmt.Scanf("%f",&F)
	C := (F-32)*(5.0/9.0)
	fmt.Print("The temperature in Celcius : ",C)
}
