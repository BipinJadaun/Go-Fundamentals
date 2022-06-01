package basics

import (
	"fmt"
)

func ArrayExp() {
	var arr1 [3]int   // integer array of size 3 decalration
	fmt.Println(arr1) // initilized with default values, for int, its 0

	arr1[0], arr1[1], arr1[2] = 1, 2, 3
	fmt.Println(arr1)

	arr2 := [3]int{4, 5, 6} // short hand decalration
	fmt.Println(arr2)

	arr3 := [3]int{12} // can declare an array with fewer elements
	fmt.Println(arr3)
	arr3[1] = 13
	arr3[2] = 15
	fmt.Println(arr3)

	// we can ommit the size in the decalration, Go will calculate the size based on values
	arr4 := [...]int{20, 22, 24, 26}
	fmt.Println(len(arr4), arr4)

	//arr4 = arr3 //can not assign a size 3 array to a different size array

	// arrays are value type. so here a copy of arr5 will be assigned to arr6. similarly in functions the arrays passed as values, not reference
	arr5 := [...]string{"India", "Russia", "USA", "Germany", "France"}

	arr6 := arr5
	arr5[0] = "Singapore" // this will not change the values in arr6
	fmt.Println("arr5 is ", arr5)
	fmt.Println("arr6 is ", arr6)

	//two dimentional array
	arr7 := [3][2]string{{"lion", "tiger"}, {"cat", "dog"}, {"pigeon", "peacock"}}
	fmt.Println(arr7)

}
