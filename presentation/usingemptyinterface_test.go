package presentation

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"testing"
)

type IDer struct {
	ID string
}

func TestEmptyInterface(t *testing.T) {
	var i interface{}
	i = 20
	fmt.Println("What is i: ", i)
	i = "hello"
	fmt.Println("What is i: ", i)
	i = struct {
		FirstName string
		LastName string
	} {"Fred", "Fredson"}
	fmt.Println("What is i: ", i)
}

func TestEmptyInterface2(t *testing.T) {
	v := struct {
		Name string `json:"full_name"`
		Email string `json:"email_default"`
		Rate float64 `json:"known_rate"`
	}{}
	str := `
{
	"full_name": "Bilal",
	"email_default": "bilal@example.com",
	"known_rate": 0.8
}
`
	err := json.Unmarshal([]byte(str), &v)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Log(v)
}

func TestEmptyInterface3(t *testing.T) {
	var data []IDer

	contents, err := ioutil.ReadFile("testdata/sample.json")
	t.Log(string(contents))
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(contents, &data)

	t.Log(data)
}

func TestEmtpyInterface4(t *testing.T) {
	var i interface{}
	i = 20
	_ = IType(i)
}

func IType(i interface{}) error {
	i2, ok := i.(string)
	if !ok {
		return errors.New("format error, expecting string input")
	}
	fmt.Println(i2)
	return nil
}