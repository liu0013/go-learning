package main

import "fmt"

func main() {
	const (
		USD int = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{EUR: "#", RMB: "&"}
	fmt.Println(RMB, symbol[RMB], len(symbol))
}