package basics

import "fmt"

//A map is a builtin type in Go that is used to store key-value pairs.
//Similar to slices, maps are reference types. So cant be compared using "==" operator.
//When a map is assigned to a new variable, they both point to the same internal data structure. So changes made in one will reflect in other.
func MapExp() {

	//A map can be created by passing the type of key and value to the make function.
	employeeSalary := make(map[string]int)

	employeeSalary["Mike"] = 1000
	employeeSalary["Steve"] = 1200
	employeeSalary["John"] = 800

	fmt.Println(employeeSalary)

	//short hand declaration
	employeeSalary1 := map[string]int{"Mike": 1000, "John": 1500}
	employeeSalary1["Adam"] = 1200

	fmt.Println(employeeSalary1)

	//iterating through map using range
	for key, value := range employeeSalary {
		fmt.Println(key, value)
	}

	//Retrieving values from map. If a value is not present, it will return 0
	johnSalary := employeeSalary["John"]
	jemmySalary := employeeSalary["Jemmy"]
	fmt.Println(johnSalary, jemmySalary)

	//To check a key-value exisit in map
	jemmySalary1, isPreset := employeeSalary["Jemmy"]

	if isPreset {
		fmt.Println(jemmySalary1)
	} else {
		fmt.Println("Jemmy salary not found")
	}

	//we can delete a key-value from map using delete function
	delete(employeeSalary, "John")
	fmt.Println(employeeSalary)

}
