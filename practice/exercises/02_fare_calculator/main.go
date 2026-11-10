// Project: island fare calculator
//A small program that quotes a domestic flight fare and checks whether the customer can pay. The prices are made up.

package main

import (
	"fmt"
	"unicode/utf8"
)

const maxNameLength = 20
const childDiscount = 0.25
const infantFare = 15
const freeBaggageKg = 20
const excessBaggageRate = 2.5
const memberDiscount = 0.05

func baseFare(route string) float64 {
	switch route {
	case "MLE-FVM":
		return 120
	case "MLE-GAN":
		return 110
	case "MLE-KDO":
		return 90
	default:
		return 0
	}
}

func passengerFare(base float64, passengerType string) float64 {
	switch passengerType {
	case "child":
		return base * (1 - childDiscount)
	case "infant":
		return infantFare
	default:
		return base
	}
}

func baggageFee(kg int) float64 {
	if kg <= freeBaggageKg {
		return 0
	}
	extra := kg - freeBaggageKg
	return float64(extra) * excessBaggageRate
}

func main() {
	passengerName := "Saadhin Nafiz"
	route := "MLE-FVM"
	passengerType := "adult"
	baggageKg := 25
	isMember := true
	walletBalance := 200.0

	// 1. Name check
	nameLen := utf8.RuneCountInString(passengerName)
	if nameLen > maxNameLength {
		fmt.Printf("Name is too long (%d characters, max %d)\n", nameLen, maxNameLength)
		return
	}
	
	// 2. Route check

	base := baseFare(route)
	if base == 0 {
		fmt.Printf("Unknown route: %s\n", route)
		return
	}

	// 3. Fare, baggage and total
	fare := passengerFare(base, passengerType)
	bag := baggageFee(baggageKg)
	total := fare + bag

	// 4. Member discount
	member := "no"
	if isMember {
		total *= 1 - memberDiscount
		member = "yes"
	}

	// 5. Receipt
	fmt.Printf("%-12s%s\n", "Passenger:", passengerName)
	fmt.Printf("%-12s%s (%s)\n", "Route:", route, passengerType)
	fmt.Printf("%-12s$%8.2f\n", "Fare:", fare)
	fmt.Printf("%-12s$%8.2f\n", "Baggage:", bag)
	fmt.Printf("%-12s%s\n", "Member:", member)
	fmt.Printf("%-12s$%8.2f\n", "Total:", total)

	// 6. Payment
	var summary string
	if walletBalance >= total {
		walletBalance -= total
		summary = fmt.Sprintf("Booking confirmed. Remaining balance: $%.2f", walletBalance)
	} else {
		summary = fmt.Sprintf("Booking failed. You need $%.2f more.", total-walletBalance)
	}
	fmt.Println(summary)
}

