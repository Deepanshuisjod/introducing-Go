package main
import "fmt"

func main(){
	fmt.Print("Enter the height in feets : ") 
	var feet float32
	fmt.Scanf("%f",&feet)
	meter := feet*0.3048
	fmt.Println("The height in meters :",meter,"m")
}
