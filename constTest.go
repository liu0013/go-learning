package main

import "fmt"

const (
	_ = iota - 1 // no type const
	Sunday
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// const (
//  int type const
// 	Sunday int = 1
// 	Monday
// 	Tuesday
// 	Wednesday
// 	Thursday
// 	Friday
// 	Saturday
// )

func main() {
	fmt.Printf("Today is %v\n", Sunday)
	fmt.Printf("Today is %v\n", Tuesday)
}
