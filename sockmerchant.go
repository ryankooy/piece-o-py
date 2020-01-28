// This is only a test. Please be advised . . .

package main

import (
	"fmt"
	"os"
)

func sockmerchant(a int, b[]int) int {
	var h[] int
	var c int = 0
	for _, i := range b {
		if i == h[i] {
			h[i] += 1
		}
		h[i] = 1
	}
	for _, j := range h {
		c += h[j] / 2
	}
	return c
}

func main() {
	var socknumber = os.Args[3]
	// TODO: Take in multiple args.
	// var therest[] int os.Args[4:]
	fmt.Println("Pairs:", sockmerchant(socknumber, therest))
}