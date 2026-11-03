package main

import "fmt"

func main() {
	name1 := "Coffee"
	qty1 := 2
	price1 := 4.50

	name2 := "Tea"
	qty2 := 1
	price2 := 3.00

	name3 := "Cake"
	qty3 := 3
	price3 := 6.25

	fmt.Printf("%-12s %3d %8.2f\n", name1, qty1, price1)
	fmt.Printf("%-12s %3d %8.2f\n", name2, qty2, price2)
	fmt.Printf("%-12s %3d %8.2f\n", name3, qty3, price3)

	total := float64(qty1)*price1 + float64(qty2)*price2 + float64(qty3)*price3
	fmt.Printf("%-12s %3s %8.2f\n", "TOTAL", "", total)
}