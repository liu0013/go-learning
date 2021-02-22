package main

import (
	"bytes"
	"fmt"
)

func intsToString(array []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range array {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteRune(']')
	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3}))
}
