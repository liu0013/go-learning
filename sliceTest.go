package main

import "fmt"

func main() {
	months := [...]string{
		1: "Jan", 2: "Feb", 3: "Mar",
		4: "Apr", 5: "May", 6: "Jun",
		7: "Jul", 8: "Aug", 9: "Sep",
		10: "Orc", 11: "Num", 12: "Dec"}

	fmt.Println(months[4:])
	test := append(months[1:], "append test")
	fmt.Println(test)
	fmt.Println(cap(test))
	fmt.Println(append(months[:4], "push"))
	fmt.Println(months)
}
