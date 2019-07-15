package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println("a, b:", a, b)
	fmt.Printf("%T\n", b)
	//use * to read val

	fmt.Println(*b)
	fmt.Println(*&a)
	//chnage value of pointer

	*b = 10
	fmt.Println(a)
}
