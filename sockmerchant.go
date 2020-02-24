// This is only a test. Please be advised . . .
package main

import (
	"fmt"
	"math/rand"
)

func sockmerchant() int {
	cl := [9]int{rand.Intn(9), rand.Intn(9), rand.Intn(9), rand.Intn(9), rand.Intn(9), rand.Intn(9), rand.Intn(9), rand.Intn(9), rand.Intn(9)}
	var c int = 0
	var h []int = append(c, rand.Intn(110))
	for _, i := range cl {
		if i == h[i] {
			h[i]++
		}
		h[i] = 1
	}
	for _, j := range h {
		c += h[j] / 2
	}
	return c
}

func main() {
	// TODO: Take in multiple args.
	// var sm = os.Args[3]
	// if sockN, err := strconv.Atoi(os.Args[1]); err == nil {
	fmt.Println("Pairs:", sockmerchant())
	// }
}
