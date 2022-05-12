package main

import (
	"exportedapp/staff"
	"fmt"
	"log"
)

var employees = []staff.Employee{
	{FirstName: "John", LastName: "Smith", Salary: 30000, FullTime: true},
	{FirstName: "Sally", LastName: "Jones", Salary: 50000, FullTime: true},
	{FirstName: "Mark", LastName: "Smithers", Salary: 60000, FullTime: true},
	{FirstName: "Allan", LastName: "Anderson", Salary: 15000, FullTime: false},
	{FirstName: "Margaret", LastName: "Carter", Salary: 100000, FullTime: false},
}

func main() {
	myStaff := staff.Office{
		AllStaff: employees,
	}

	//myStaff.All()

	fmt.Println(myStaff.All())

	// staff.OverPaidLimit = 10000
	log.Println("Overpaid staff", myStaff.OverPaid())
	log.Println("Underpaid staff", myStaff.UnderPaid())

}
