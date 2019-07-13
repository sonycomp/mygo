package main

import (
	"fmt"
)

func main() {
	i := 1
	for ; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even no:", i)
		}
	}

	for j := 1; j <= 10; j++ {
		if j%2 != 0 {
			fmt.Println("Odd nos:", j)
		}
	}

}
