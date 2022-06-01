package database

import (
	"Go-Fundamentals/src/webapp/entity"
	"context"
	"fmt"
	"testing"
)

var db = MongoDB{}

func TestConnection(t *testing.T) {

	conn, _ := db.GetConn(context.TODO())
	actual := conn.Client.Database(conn.Config.DBName)
	expected := "target"
	fmt.Printf("expected %v, actual %v", expected, actual)
}

func TestInsert(t *testing.T) {
	conn, _ := db.GetConn(context.TODO())

	emp1 := entity.Employee{
		Name:        "Ajay",
		Department:  "Engineering",
		Designation: "Lead Engineer",
		Age:         34,
	}
	emp2 := entity.Employee{
		Name:        "Bipin",
		Department:  "Engineering",
		Designation: "Senior Engineer",
		Age:         30,
		Address:     entity.Address{Area: "Dayalbagh", City: "Agra", Pin: 282005},
	}
	employee := []interface{}{emp1, emp2}

	req := entity.DBRequest{Data: employee, Collection: "employee"}

	db.Insert(context.TODO(), *conn, req)

}
