package staff

import "log"

var OverPaidLimit = 75000
var underPaidLimit = 20000

type Employee struct {
	FirstName string
	LastName  string
	Salary    int
	FullTime  bool
}

type Office struct {
	AllStaff []Employee
}

func (e *Office) All() []Employee {
	return e.AllStaff
}

func (e *Office) OverPaid() []Employee {
	var overpaid []Employee

	for _, emp := range e.AllStaff {
		if emp.Salary > OverPaidLimit {
			overpaid = append(overpaid, emp)
		}
	}
	return overpaid
}

func (e *Office) UnderPaid() []Employee {
	var underpaid []Employee
	myFunction()

	for _, emp := range e.AllStaff {
		if emp.Salary < underPaidLimit {
			underpaid = append(underpaid, emp)
		}
	}
	return underpaid
}

func (e *Office) notVisible() {
	log.Println("Hello World")
}

func myFunction() {
	log.Println("I am a function")
}
