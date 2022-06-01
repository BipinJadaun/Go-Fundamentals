package basics

import "fmt"

// struct is similar to a class in java, where we can declare variables and methods
// struct name and field names starts with capital letter, will be accessible outside the package
// fields will be initialized with zero(default) values.
type Employee struct {
	name        string
	age         int
	department  string
	designation string
	address     Address //nested struct
}

type Address struct {
	area string
	city string
	pin  int
}

func StructExp() {
	//creating struct with field names. fields can be in different order.
	// we can create struct with few values
	emp1 := Employee{
		name:        "Bipin",
		department:  "Engineering",
		designation: "Senior Engineer",
		age:         30,
	}
	fmt.Println(emp1)

	//creating struct with values only. values must be in same order as decalred in struct
	emp2 := Employee{"Bipin", 30, "Engineering", "Senior Engineer", Address{"Dayalbagh", "Agra", 282005}}
	fmt.Println(emp2)

	//the individual fields from struct can accessed using "." operator
	emp1.address = Address{"Dayalbagh", "Agra", 282005}
	fmt.Println(emp1.name, emp1.address)

	//Structs are value types and are comparable if each of their fields are comparable.
	fmt.Println(emp1 == emp2)
	emp1.age = 31
	fmt.Println(emp1 == emp2)

}
