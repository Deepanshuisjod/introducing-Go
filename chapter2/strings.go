package main
import "fmt"

func main(){
	name := "Deepanshu"
	first_char := string(name[0]) // This will return the actual first character
	first_char_ASCII := name[0] // This will return the ASCII of the first character
	
	fmt.Println(len("Hello World")) // Will print out the length of the given string
	fmt.Println("Hello, World"[1]) // Will give the ASCII value of first character i.e 'e'
	fmt.Println("Hello, " + "World") // Will concat the two strings
	
	fmt.Println(first_char)
	fmt.Println(first_char_ASCII)
}

