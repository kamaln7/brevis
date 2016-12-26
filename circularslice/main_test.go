package circularslice

import (
	"testing"
)

func strify(s []interface{}) string {
	var res string
	for _, c := range s {
		res += c.(string)
	}

	return res
}

func TestCircularProperty(t *testing.T) {
	slice := New(4)
	slice.Insert("h").Insert("e").Insert("l").Insert("l").Insert("o")
	value := strify(slice.Get())
	if value != "ello" {
		t.Fail()
	}
	slice.Insert("t").Insert("5")
	value = strify(slice.Get())
	if value != "lot5" {
		t.Fail()
	}
}
