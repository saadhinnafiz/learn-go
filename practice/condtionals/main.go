// if statements in Go do not use parentheses around the condition:
package main

import "fmt"

func main() {
	height := 5
	if height > 4 {
	    fmt.Println("You are tall enough!")
	}
}
