package main

import "fmt"

func main() {
	total := 300.0

	if total <= 100 {
		total *= 0.9
	} else if total > 100 && total <= 300 {
		total *= 0.8
	} else if total > 300 {
		total *= 0.5
	}
	fmt.Printf("total: %.2f\n", total)
	total = 400
	switch {
	case total <= 100:
		total *= 0.9
	case total > 100 && total <= 300:
		total *= 0.8
	case total > 300:
		total *= 0.5
	default:
		total = 2
	}
	fmt.Printf("total: %v\n", total)

}
