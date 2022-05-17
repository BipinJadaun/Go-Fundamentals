package learning

import "fmt"

//A slice is a convenient, flexible and powerful wrapper on top of an array.
//Slices do not own any data on their own. They are just references to existing arrays.

func SliceExp() {
	arr := [6]int{1, 2, 3, 4, 5, 6}

	slice1 := arr[1:4] // creates a slice from 1 to 4 (start: included, end: excluded)
	fmt.Println("slice1:", slice1)

	slice2 := []int{6, 7, 8} //creates and array and returns a slice reference
	fmt.Println("slice2:", slice2)

	slice3 := arr[:] //creates a slice which contains all elements of the array
	fmt.Println("slice3:", slice3)

	//The length of a slice is the number of elements in the slice.
	//The capacity of a slice is the number of elements in the underlying array starting from the index from which the slice is created.
	fmt.Println("length of slice1:", len(slice1), "and capacity is:", cap(slice1))

	// We can create a slice variable using make function by passing the type, length and capacity.
	// The capacity parameter is optional and defaults to the length.
	slice4 := make([]int, 4, 7)
	slice4[0], slice4[1] = 1, 2
	fmt.Println("slice4:", slice4)

	//length of a slice can be modified upto its capacity
	slice1 = slice1[:cap(slice1)]
	fmt.Println("length of slice1:", len(slice1), "and capacity is:", cap(slice1))

	//arrays are restricted to fixed length and their length cannot be increased.
	//But Slices are dynamic and new elements can be appended to the slice using "append" function.
	//Internally a new array gets created. The elements of the existing array are copied to this new array.
	//And a new slice reference for this new array is returned. The capacity of the new slice is now twice that of the old slice.

	fruits := []string{"apple", "banana", "orange"}
	fmt.Println("length of fruit slice:", len(fruits), "and capacity is:", cap(fruits))
	fruits = append(fruits, "mango")
	fmt.Println("length of fruit slice:", len(fruits), "and capacity is:", cap(fruits))

	//we can also append one slice to another of same type using "..." variadic operator
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	foods := append(fruits, veggies...)

	fmt.Println(foods)
	fmt.Println("length of food slice:", len(foods), "and capacity is:", cap(foods))

	//we can create a copy of slice using copy function
	//this is helful to optimize memory since the original array can now be garbage collected
	onlyFruits := foods[:len(foods)-3]
	fruitsCopy := make([]string, len(onlyFruits))
	copy(fruitsCopy, onlyFruits) //copies onlyFruits to fruitsCopy
	fmt.Println(fruitsCopy)

}
