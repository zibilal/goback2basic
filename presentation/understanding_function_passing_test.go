package presentation

import (
	"sort"
	"testing"
)

type Person struct {
	FirstName string
	LastName string
	Age int
}

func TestFunctionPassing(t *testing.T) {
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}
	t.Log("People", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	t.Log("People", people)
}
