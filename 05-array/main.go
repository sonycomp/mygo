package main

import "fmt"

func main() {
	//Array
	var fruitarray [2]string
	//Add values
	fruitarray[0] = "Apple"
	fruitarray[1] = "Mango"

	fmt.Println(fruitarray)
	newarr := [2]string{"apple", "orange"}
	fmt.Println(newarr)

	fslice := []string{"orange", "mango", "banana"}
	fmt.Println(fslice[1:2])

}
