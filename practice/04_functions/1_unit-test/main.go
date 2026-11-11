/*

Assignment

Complete the getMonthlyPrice function. It accepts a tier (string) as input and returns the monthly price for that tier in pennies. Here are the prices in dollars:

    "basic" - $100.00
    "premium" - $150.00
    "enterprise" - $500.00

Convert the prices from dollars to pennies. If the given tier doesn't match any of the above, return 0 pennies.

*/

package main

import "fmt"

func getMonthlyPrice(tier string) int {
	switch tier {
	case "basic":
		return 100 * 100
	case "premium":
		return 150 * 100
	case "enterprise":
		return 500 * 100
	default:
		return 0
	}
}

func main() {
	fmt.Println(getMonthlyPrice("basic"))
	fmt.Println(getMonthlyPrice("premium"))
	fmt.Println(getMonthlyPrice("enterprise"))
	fmt.Println(getMonthlyPrice("free"))
}