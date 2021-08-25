package presentation

import (
	"testing"
)

func TestMapNilRead(t *testing.T){
	var nilMap map[string]int

	t.Log(nilMap)
}

func TestMapNilWrite(t *testing.T){
	defer func() {
		if recover() != nil {
			t.Log("PANIC!!")
		}
	}()

	var nilMap map[string]int

	nilMap["ok"] = 15

	t.Log(nilMap)
}

func TestInitWithMake(t *testing.T){

	amap := make(map[string]int)
	amap["ok"] = 15

	t.Log(amap)
}

func TestMapWithOkIdiom(t *testing.T) {
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	t.Log(v, ok)

	v, ok = m["world"]
	t.Log(v, ok)

	v, ok = m["goodbye"]
	t.Log(v, ok)
}

func TestMapDeleteValuie(t *testing.T) {

	m := map[string]int{
		"hello": 5,
		"world": 10,
	}
	t.Log("Before", m)
	delete(m, "hello")
	t.Log("Afater", m)
}

func TestMapAsSet(t *testing.T) {
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	t.Log(len(vals), len(intSet))
	t.Log(intSet)
	t.Log(intSet[5])
	t.Log(intSet[500])
	if intSet[100] {
		t.Log("100 is in the set")
	}
}
