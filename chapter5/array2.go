package main
import "fmt"

func main(){
	var arr[5] int8
	var total int8 = 0
	for i:= 0; i < len(arr); i++{
		fmt.Print("Enter the element at ",i,"position : ")
		fmt.Scanf("%d",&arr[i])
		total += arr[i]
	}
	fmt.Println("The array is : ",arr)
	// Finding the average of the array
	fmt.Print("The average of the array is : ")
	fmt.Println(total/int8(len(arr)))
	
	// Slicing 
	sliced_arr := arr[0:3]
	fmt.Println(sliced_arr)
}


