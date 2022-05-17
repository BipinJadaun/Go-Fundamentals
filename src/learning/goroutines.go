package learning

import (
	"fmt"
	"time"
)

// Goroutines are functions or methods that run concurrently with other functions or methods.
// Goroutines are lightweight threads managed by Go runtime as compared to java where treads are OS threads and managed by CPU scheduler.
// Goroutines are extremely cheap when compared to threads and the stack can grow and shrink according to the needs.

func GoroutineExp() {

	go numbers() // call the fun with 'go' keyword to run the function as goroutine
	go alphabets()

	//The main Goroutine should be running for any other Goroutines to run.
	//If the main Goroutine terminates then the program will be terminated and no other Goroutine will run.
	time.Sleep(1000 * time.Millisecond)

	fmt.Println("main terminated")

}

func numbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("%c ", i)
	}
	fmt.Println()
}
