package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Printf("i: %v\t", i)
	}

	fmt.Println()

	j := 0
	for j < 4 {
		fmt.Printf("j: %v\t", j)
		j++
	}

	fmt.Println()

	t := 0
	for {
		if t >= 4 {
			break
		}
		fmt.Printf("t: %v\t", t)
		t += 1
	}

}
