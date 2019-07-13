package main

import "fmt"

func greet(name string) string {
	return "Hello " + name
}

func getfullname(first, second string) string {
	return first + " " + second
}

func multi(n1, n2, n3 int) int {
	return n1 * n2 * n3
}
func main() {
	fmt.Println(greet("Rahul"))
	fmt.Println(getfullname("Rahul", "Bhandari"))
	fmt.Println(multi(2, 3, 2))
}
