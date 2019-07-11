package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("Name: %s - Age: %d\n", p.Name, p.Age)
}

type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

type ByName []Person

func (a ByName) Len() int {
	return len(a)
}

func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

func main() {
	p1 := Person{
		Name: "Donald Duck",
		Age:  80,
	}

	p2 := Person{
		Name: "Mickey Mouse",
		Age:  90,
	}

	p3 := Person{
		Name: "Goofy",
		Age:  75,
	}

	people := []Person{p1, p2, p3}

	fmt.Println("Unsorted: ", people)

	sort.Sort(ByAge(people))
	fmt.Println("Sorted by age:\n", people)

	sort.Sort(ByName(people))
	fmt.Println("Sorted by name:\n", people)
}
