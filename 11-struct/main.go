package main

import "fmt"

//Person isexported
type Person struct {
	firstname string
	lastname  string
	age       int
	zip       int
	blind     bool
}

func main() {
	//init struct

	//	person1 := Person{firstname: "Rahul", lastname: "Bhandari", age: 43, zip: 23, blind: false}

	person1 := Person{"Rahul", "Bhandari", 43, 23, false}

	fmt.Println(person1)

	//Print only age

	fmt.Println(person1.age)

}
