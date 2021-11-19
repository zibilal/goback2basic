package presentation

import (
	"fmt"
	"testing"
)

type Enablement struct {
	SquadNames []string
	Chapters []string
	Total int
}


func TestUsingStruct(t *testing.T) {

	var i int
	fmt.Println(i)

	enablement1 := Enablement{
		Total: 15,
	}

	enablement2 := struct {
		SquadNames []string
		Chapters []string
		Total int
	}{
		Total: 60,
		Chapters: []string{"Mobile Banking", "Internet Banking"},
		SquadNames: []string{"Mobile Banking", "Internet Banking"},
	}

	fmt.Println("Enablement2", enablement2, "Enablement1", enablement1.Chapters)

	enablementOneMap := map[string][]string{
		"SquadNames": {"Mobile Banking", "Internet Banking"},
		"Chapters": {"Mobile Banking", "Internet Banking"},
	}

	fmt.Println(enablementOneMap)

	enablementSatu := Enablement{}
	enablementDua := struct {
		SquadNames []string
		Chapters []string
		Total int
	}{}

	fmt.Println("Enablement Satu", enablementSatu)
	fmt.Println("Enablement Dua", enablementDua)
	enablementSatu = enablementDua

	var satuan int8
	duaan := int16(satuan)
	print(duaan)

	{
		satuan = 15
		print(satuan)
	}

	print(satuan)
}
