package learning

import "fmt"

//global valriable accessible across the functions
var topic string // we can have unused global variables

func VariableExp() {

	topic = "Go Fundamentals"
	fmt.Println(topic)

	var shape1 string //Declaration

	shape1 = "Cube" //assignment, values must be of type used in declaration

	var length = 2 //declaration and assignment. Here type is not mandatory, Go will infer the Type automatically from value.

	var width, hight = 3, 4 // we can declare multiple variables. but number of values must be same as number of variables.

	volume := length * width * hight //Short hand declaration

	//count := 1              // unused variables in funtions are not supported

	fmt.Println("shape:", shape1, "volume:", volume)

}
