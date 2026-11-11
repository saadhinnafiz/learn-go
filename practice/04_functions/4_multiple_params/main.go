package main

import "fmt"

func describe(hp, damage int, name string, level int) string {
	return fmt.Sprintf("%s (level %d): hp=%d, damage=%d", name, level, hp, damage)
}

func hitDamage(damage, armor int, multiplier float64) float64 {
	dealt := damage - armor
	if dealt < 0 {
		dealt = 0
	}
	return  float64(dealt)*multiplier
}

func main() {
	fmt.Println(describe(40, 7, "Goblin", 3))
	fmt.Println(describe(250, 30, "Dragon", 12))
	fmt.Printf("%.1f\n", hitDamage(20, 5, 1.5))
	fmt.Printf("%.1f\n", hitDamage(10, 12, 2.0))
	fmt.Printf("%.1f\n", hitDamage(30, 0, 1.0))
}

