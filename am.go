package main

import (
	"fmt"
)

func main() {
	g := G()
	if g == nil {
		fmt.Println(11)
	}
}

func G() []int {
	return nil
}
