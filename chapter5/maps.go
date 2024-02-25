package main
import "fmt"

func main(){
	// The map is like dictionary in python
	// Map has {key:value} 
	// You can't use maps before declaring it
	x := make(map[string]int) // Maps have to be initiazed before they can be used
	x["id1"] = 10
	x["id2"] = 11
	x["id3"] = 12
	fmt.Println(x)
}
