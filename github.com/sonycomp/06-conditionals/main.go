package main

import "fmt"

func main() {
	x := 5
	y := 4
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else if x > y {
		fmt.Printf("%d is grater than %d\n", x, y)
	} else {
		fmt.Printf("%d is equals to %d\n", x, y)
	}

	//Fiz Buz

	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Printf("%d FizzBuzz\n", i)
		} else if i%3 == 0 {
			fmt.Printf("%d Fizz\n", i)
		} else if i%5 == 0 {
			fmt.Printf("%d Buzz\n", i)
		} else {
			fmt.Println(i)
		}
	}

}
