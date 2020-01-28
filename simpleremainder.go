package main

import (
	"fmt"
	"math"
)

func rem(a, b float64) float64 {
	return math.Remainder(a, b)
}

func main() {
	fmt.Println("The sum is actually", rem(93486062398, 2395029302), ", ya doofus! Love you. :)")
}