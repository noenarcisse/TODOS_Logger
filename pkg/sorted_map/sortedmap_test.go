package sortedmap

import (
	"slices"
	"testing"
)

func Test_Add(t *testing.T) {
	sm := New[int, struct{}]()
	sm.Add(10, struct{}{})

	if len(sm.inner) != len(sm.order) {
		t.Errorf(`Add()
		inner and order len are different
		inner = %d
		order = %d`, len(sm.inner), len(sm.order))
	}

	if _, ok := sm.inner[10]; !ok {
		t.Errorf(`Add()
		the key added is not present in SortedMap.inner`)
	}
	if !slices.Contains(sm.order, 10) {
		t.Errorf(`Add()
		the key added is not present in SortedMap.order`)
	}
}

func Test_Add2(t *testing.T) {
	sm := New[int, struct{}]()
	sm.Add(10, struct{}{})
	sm.Add(10, struct{}{})

	count := len(sm.order)

	if count != 1 {
		t.Errorf(`Add()
		the number of keys found in the order is incorrect
		expected: 1
		found: %d
		`, count)
	}
}

func Test_Clear(t *testing.T) {
	sm := New[int, string]()
	sm.Add(1, "Un")
	sm.Add(2, "Deux")
	sm.Add(3, "Trois")

	sm.Clear()

	exp := 0
	res1 := len(sm.inner)
	res2 := len(sm.order)

	if res1 > 0 || res2 > 0 {
		t.Errorf(`Clear(). 
		Result: %d, %d
		Expected: %d, %d`,
			res1, res2, exp, exp)
	}
}
