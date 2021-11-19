package presentation

import (
	"fmt"
	"testing"
	"time"
)

func TestUsingMap(t *testing.T) {
	httpStatus := map[string]int{}

	httpStatus["Not Found"] = 404
	httpStatus["Bad Request"] = 400

	var squad = make(map[string][]string)
	squad["lombok"] = []string{
		"Arifullah", "Arief Permana", "Marcel",
	}
	squad["komodo"] = []string{ "Nanda", "Stephen", "Kadek" }

	for i, v := range(squad["komodo"]) {
		fmt.Println(i, v)
	}

	httpStatusCopy := map[string]int {
		"Not Found": 404,
		"Bad Request": 400,
	}

	fmt.Println(httpStatusCopy)

	lombokService := map[string]time.Time{}
	lombokService["fin-prs"] = time.Now().Add(-5 * 24 * time.Hour)
	lombokService["funded-things"] = time.Now()

	fmt.Println(lombokService)
	fmt.Println("----------------------")
	for k, v := range lombokService {
		fmt.Println(k, v)
	}

	delete(lombokService, "funded-things")
	fmt.Println("----------------------")
	for k, v := range lombokService {
		fmt.Println(k, v)
	}

	fmt.Println("Disini", lombokService["funded-things"])

	v, found := lombokService["funded-things"]
	if found {
		fmt.Println("Found!", v)
	} else {
		fmt.Println("Could not found funded-things")
	}


}
