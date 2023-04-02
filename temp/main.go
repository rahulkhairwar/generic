package main

import "fmt"

func main() {

}

func iterableSlices() {
	var s []int
	var p []person
	var a []address

	s = []int{1, 2, 3, 4, 5}
	a = []address{
		{
			street: "123 Main St",
			city:   "New York",
			pin:    12345,
			county: "USA",
		},
		{
			street: "456 Main St",
			city:   "New York",
			pin:    12345,
			county: "USA",
		},
	}
	p = []person{
		{
			name:    "John",
			age:     30,
			address: a[0],
		},
		{
			name:    "Jane",
			age:     30,
			address: a[1],
		},
	}

	forEach(p, func(v int) {
		fmt.Println(v)
	})
}

type person struct {
	name    string
	age     int
	address address
}

type address struct {
	street string
	city   string
	pin    int
	county string
}

type iterable[T any] []T

func forEach[T any](iter iterable[[]any], fn func(T)) {
	for _, v := range iter {
		fmt.Println(v)
	}
}
