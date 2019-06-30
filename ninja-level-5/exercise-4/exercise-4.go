package main

import "fmt"

func main() {
	avgGrades := struct {
		classroom string
		students  map[string]int
		subjects  []string
	}{
		classroom: "A",
		students: map[string]int{
			"Al":   2,
			"John": 7,
			"Jack": 5,
		},
		subjects: []string{
			"PE",
			"Math",
			"History",
		},
	}

	fmt.Println(avgGrades.classroom)
	fmt.Println(avgGrades.students)
	fmt.Println(avgGrades.subjects)

	for k, v := range avgGrades.students {
		fmt.Println(k, v)
	}

	for i, v := range avgGrades.subjects {
		fmt.Println(i, v)
	}
}
