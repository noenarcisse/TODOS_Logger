package sortedmap

import (
	"cmp"
	"errors"
	"iter"
	"maps"
	"slices"
)

type SortedMap[T cmp.Ordered, U any] struct {
	inner  map[T]U
	order  []T
	sorted bool
}

func New[T cmp.Ordered, U any]() *SortedMap[T, U] {
	return &SortedMap[T, U]{
		inner:  map[T]U{},
		order:  []T{},
		sorted: true, //optimistic with Add trying to pass this to false whenever possible
	}
}

func FromMap[T cmp.Ordered, U any](m map[T]U) *SortedMap[T, U] {
	sm := New[T, U]()
	sm.order = slices.Collect(maps.Keys(m))
	sm.inner = m
	return sm
}

func Zip[T cmp.Ordered, U any](t []T, u []U) (*SortedMap[T, U], error) {

	if len(t) != len(u) {
		return nil, errors.New("The two elements have different lengths. Couldn't zip.")
	}

	sm := New[T, U]()
	for i := range t {
		if _, ok := sm.Get(t[i]); ok {
			return nil, errors.New("The first list contains doubles. Couldn't zip")
		}
		sm.Add(t[i], u[i])
	}

	return sm, nil
}

func (s *SortedMap[T, U]) Clear() {
	s.inner = map[T]U{}
	s.order = []T{}
}

func (s *SortedMap[T, U]) sort() {
	slices.Sort(s.order)
	s.sorted = true
}

func (s *SortedMap[T, U]) Add(key T, val U) {

	if s.inner == nil {
		s.inner = make(map[T]U)
		s.order = make([]T, 0)
		s.sorted = true
	}

	if _, ok := s.inner[key]; !ok {
		if len(s.order) > 0 && key < s.order[len(s.order)-1] {
			s.sorted = false
		}
		s.order = append(s.order, key)
	}
	s.inner[key] = val
}

func (s *SortedMap[T, U]) Get(key T) (U, bool) {
	val, ok := s.inner[key]
	return val, ok
}

func (s *SortedMap[T, U]) Len() int {
	return len(s.inner)
}

// naming conv go -> All() instead of Items?

// doc
func (s *SortedMap[T, U]) Items() iter.Seq2[T, U] {

	if !s.sorted {
		s.sort()
	}

	return func(yield func(T, U) bool) {
		for _, e := range s.order {
			if !yield(e, s.inner[e]) {
				return
			}
		}
	}
}
