/*
Exercise: Shipping Cost Calculator

Write a function getShippingCost that accepts a speed (string) and returns the shipping cost in cents (an int).

Prices:

"standard" — $5.00
"express" — $15.00
"overnight" — $40.00
Anything else — return -1 (to signal an invalid/unrecognized speed)

Try writing the full function signature and body yourself, in your local learn-go project, without looking back at the last example.

Once you've got a draft, add a temporary main() to test it manually:

*/

package main

import "fmt"


func getShippingCost(speed string) int {
	switch speed {
	case "standard":
		return 100 * 5 // $5.00 in cents
	case "express":
		return 100 * 15 // $15.00 in cents
	case "overnight":
		return 100 * 40 // $40.00 in cents
	default:
		return -1 // Invalid/unrecognized speed
	}
}

func main() {
	fmt.Println(getShippingCost("standard"))  // Output: 500
	fmt.Println(getShippingCost("express"))   // Output: 1500
	fmt.Println(getShippingCost("overnight")) // Output: 4000
	fmt.Println(getShippingCost("invalid"))   // Output: -1
}