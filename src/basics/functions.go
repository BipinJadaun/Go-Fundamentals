package basics

import (
	"fmt"
)

func FunctionExp() {

	length, width, hight := 2, 4, 6

	parimeter := calculateParimeter(length, width)
	fmt.Println("parimeter: ", parimeter)

	area2, volume2 := caclcualteAreaAndVolume(length, width, hight)
	fmt.Println("parimeter: ", parimeter, "area2: ", area2, "volume2: ", volume2)

	// we can assign only required values and discard rest of the values returned by function
	area1, _ := caclcualteAreaAndVolume1(length, width, hight)
	fmt.Println("area1: ", area1)
}

// function with single return value
func calculateParimeter(length int, width int) int {
	variadicFunctionExp(length, width)
	return 2 * (length + width)
}

// function with multiple return values
func caclcualteAreaAndVolume(length, width, hight int) (int, int) {
	variadicFunctionExp(length, width, hight)
	area := length * width
	volume := length * width * hight
	return area, volume
}

//function with named return values. Here we dont have to mention the returnig values with return statement
func caclcualteAreaAndVolume1(length, width, hight int) (area, volume int) {
	area = length * width
	volume = length * width * hight
	return
}

//variadic functions can be invoked with any number of similar type arguments
func variadicFunctionExp(dimentions ...int) {
	fmt.Println("number of args:", len(dimentions))
}
