package main

import "fmt"

func main() {
	fmt.Println(RemainingOvenTime(5))
}

const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}
