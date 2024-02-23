/*
Go supports the following standard arithmetic operators:
+ -> Addition
- -> Subtraction
* -> Multiplication
/ -> Division
% -> Remainder
*/
package main

import "fmt"

func main(){
	fmt.Println("1 + 1 =",1 + 1) // Adding two integers using '+' operator
	fmt.Println("1 + 1 =",1.0 + 1.0) // Adding two integers using '+' operator but used floating point literal
	fmt.Println(3*3)
	fmt.Println(90/8)
	fmt.Println(90%6)
}
