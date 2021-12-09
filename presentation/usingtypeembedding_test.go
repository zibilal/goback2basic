package presentation

import (
	"fmt"
	"testing"
)

type Employee struct {
	Name     string
	ID       string
	IsSenior bool
}

func (e Employee) String() string {
	return fmt.Sprintf("Yang INI: %s (%s) %t", e.Name, e.ID, e.IsSenior)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	// do business logic
	// get data from databases
	return m.Reports
}

func (m Manager) String() string {
	return fmt.Sprintf("Yang ITU: %s (%s) %t, Reports: %d", m.Name, m.ID, m.IsSenior, len(m.Reports))
}

func TestManagerType(t *testing.T) {
	// composition --> has a relationship
	irvan := Manager{
		Employee: Employee{
			Name: "Irvan",
			ID:   "012378",
			IsSenior: true,
		},
		Reports: []Employee{
			{Name: "Fauzan", ID: "012"},
			{Name: "Chandra", ID: "013"},
		},
	}

	fmt.Println("Irvan: ", irvan)
}

func AATest(e Employee) {
	fmt.Println(e)
}
