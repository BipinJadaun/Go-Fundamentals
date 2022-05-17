package learning

import (
	"fmt"
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
