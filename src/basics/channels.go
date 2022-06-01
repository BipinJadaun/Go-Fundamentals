package basics

import (
	"fmt"
	"sync"
)

//Channels can be thought of as pipes using which Goroutines communicate. Data can be sent from one end and received from the other end using channels.

func ChannelExp() {
	// create a channel. awe can pass the size of channel as well (optional)
	sqrChannel := make(chan int)
	cubeChennel := make(chan int)

	num := 12

	go calcSqrSum(num, sqrChannel)
	go calcCubeSum(num, cubeChennel)

	sqrSum := <-sqrChannel // blocking call, main goroutine will wait untill a goroutine writes data into channel
	cubeSum := <-cubeChennel

	fmt.Println("SquareSume:", sqrSum)
	fmt.Println("CubeSume:", cubeSum)
}

func calcSqrSum(num int, sqrChannel chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum = sum + digit*digit
		num = num / 10
	}
	sqrChannel <- sum // adding data into channel
}

func calcCubeSum(num int, cubeChannel chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum = sum + digit*digit*digit
		num = num / 10
	}
	cubeChannel <- sum // adding data into channel
}

func SynchonizationUsingChannelExp() {

	// schanel with capacity 1.
	chanel := make(chan bool, 1)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&wg, chanel)
	}

	wg.Wait() //wait for wait group counter to becomes 0

	fmt.Println("final value of x:", x)
}

var x = 1 // global variable to be used by multiple goroutines

func increment(wg *sync.WaitGroup, chanel chan bool) {

	//all other Goroutines trying to write to this channel are blocked until the value is read from this channel after incrementing x
	chanel <- true

	x = x + 1

	<-chanel

	wg.Done()
}
