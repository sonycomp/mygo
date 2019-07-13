package main

import "fmt"

func main() {
	//Declare map

	mymap := make(map[string]string)

	mymap["rahul"] = "rahul@att.net"
	mymap["dhamu"] = "dhamu@yahoo.com"

	fmt.Println(mymap["dhamu"])

	//assign values

	mymap2 := map[string]int{"Age": 42, "Zip": 42, "Area": 32}
	fmt.Println(mymap2)
	delete(mymap2, "Zip")
	fmt.Println(mymap2)
}
