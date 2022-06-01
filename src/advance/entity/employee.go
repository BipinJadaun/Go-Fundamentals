package entity

type Employee struct {
	Name        string
	Age         int
	Department  string
	Designation string
	Address     Address //nested struct
}

type Address struct {
	Area string
	City string
	Pin  int
}
