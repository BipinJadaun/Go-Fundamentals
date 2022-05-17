package learning

import "fmt"

func LoopExp() {

	// initialization; condition; post operation
	for i := 1; i <= 15; i++ {
		if i%2 == 1 {
			continue // skips the current iteration
		} else if i > 10 {
			break // breaks the loop
		}
		fmt.Println(i)
	}

	i := 0        // initialization
	for i <= 10 { // initialisation and post are omitted
		fmt.Printf("%d ", i)
		i += 2
	}

	// nested loops, we can break the outer loop from the inner loop by using leble on outer for loop
outer:
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
			if i == j {
				break outer // will break the outer for loop
			}
		}
		fmt.Println()
	}

	arr := [4]int{1, 2, 3, 4}
	// using Range for array iteration
	for i, v := range arr {
		fmt.Println("index: ", i, "value: ", v)
	}

}
