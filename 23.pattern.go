package main

import (
	"fmt"
)

func MapKeys[k comparable, v any](m map[k]v) []k {
	r := make([]k, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type List[t any] struct {
	head, tail *element[t]
}
type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[t]) Push(v t) {
	if lst.tail == nil {
		lst.head = &element[t]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[t]{val: v}
		lst.tail = lst.tail.next
	}
}
func (lst *List[t]) GetAll() []t {
	var elems []t
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}
	fmt.Println("keys m:", MapKeys(m))
	_ = MapKeys[int, string](m)
	lst := List[int]{}
	lst.Push(10)
	lst.Push(20)
	lst.Push(30)
	fmt.Println(lst.GetAll())
}
