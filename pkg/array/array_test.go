package array

import (
	"reflect"
	"testing"
)

func TestFind(t *testing.T) {
	var array = []int{1, 2, 3, 4, 5}
	var result = Find(array, func(v int) bool {
		return v == 3
	})
	if *result != 3 {
		t.Error("Find failed, expected 3, got ", *result)
	}
}

type Item struct {
	Name string
}

func TestMap(t *testing.T) {
	var array = []Item{
		{Name: "test"},
		{Name: "testing"},
	}
	result := Map(array, func(v Item) string {
		return v.Name
	})
	if reflect.DeepEqual(result, []string{"test", "testing"}) == false {
		t.Error("Map failed, expected [test testing], got ", result)
	}
}

func TestFilter(t *testing.T) {
	var array = []int{1, 2, 3, 4, 5}
	result := Filter(array, func(v int) bool {
		return v > 3
	})
	if reflect.DeepEqual(result, []int{4, 5}) == false {
		t.Error("Filter failed, expected [4 5], got ", result)
	}
}
