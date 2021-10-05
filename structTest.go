package main

import (
	"aaron/go/learning/geodemo"
	"fmt"
)

func main() {
	c := new(geodemo.ColoredPoint)
	c.X = 1
	c.Y = 2
	c.ScaleBy(2)
	fmt.Printf("X: %v, Y: %v", c.X, c.Y)
}
