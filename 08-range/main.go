package main

import "fmt"

func main() {
	ids := []int{3, 5, 6, 8, 9}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}
	//no using index

	col := []int{3, 5, 6, 8, 9}

	for _, nos := range col {
		fmt.Printf("nos: %d\n", nos)

	}
	//adding all nos. from slice

	nos := []int{4, 7, 9}
	sum := 0
	for _, idx := range nos {
		sum += idx
	}
	fmt.Println(sum)

	//range with maps

	map1 := map[string]string{"rahul": "rahul@xyz.com", "soham": "soham@xyz.com", "dhamu": "dhamu@xyz.com"}

	for k, v := range map1 {
		fmt.Printf("%s:%s\n", k, v)

		//print k only

		for k := range map1 {
			fmt.Println("Name:" + k)
		}
	}
}
