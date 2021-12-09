package presentation

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

type ProsperaIDDefault uint8

func TestUsingType(t *testing.T) {
	var tmp uint8

	tmp = uint8(GenerateProsperaID(ProsperaIDDefault(30) ))
	fmt.Println("Tmp", tmp)
}

func GenerateProsperaID(latest ProsperaIDDefault) ProsperaIDDefault{
	return latest + 1
}

type Score int
type Converter func(string)Score

func Action(input string, action Converter) Score {
	return action(input)
}

func ConvertStringToScore(input string) Score{
	i, _ := strconv.Atoi(input)
	return Score(i)
}

func TestA(t *testing.T) {
	result := Action("123", ConvertStringToScore)
	fmt.Println("Result: ", result)
}

func TestB(t *testing.T) {
	result := Action( "123", func(input string) Score {
		i, _ := strconv.Atoi(input)
		return Score(i)
	})
	fmt.Println("Result: ", result)
}

type PersonEx struct {
	FirstName string
	LastName string
	Age int
}

func (p PersonEx) String() string {
	return fmt.Sprintf("FName: %s LName: %s, age %d", p.FirstName, p.LastName, p.Age)
}

func (p *PersonEx) UpdateAge(age int) *PersonEx {
	p.Age = age
	return p
}

func (p *PersonEx) UpdateFName(name string) *PersonEx {
	 p.FirstName = name
	 return p
}

func (p *PersonEx) UpdateLName(name string) *PersonEx {
	p.FirstName = name
	return p
}

func (p *PersonEx) Update(input *PersonEx) (*PersonEx, error){
	if p == nil {
		return nil, errors.New("object is empty")
	}
	p.UpdateFName(input.FirstName).UpdateLName(input.LastName).UpdateAge(input.Age)
	return p, nil
}

func TestPersonEx(t *testing.T) {
	p := PersonEx{
		"Tester", "Automation", 25,
	}
	fmt.Println(p)
	p.UpdateAge(30)
	fmt.Println(p)
	p.UpdateFName("Developer")
	p.UpdateLName("Back-end")
	fmt.Println(p)
}
