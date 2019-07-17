package main

import (
	"fmt"

	"github.com/grizzlygr/golang-course-exercises/ninja-level-13/exercise-2/quote"
	"github.com/grizzlygr/golang-course-exercises/ninja-level-13/exercise-2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		//fmt.Printf("test{\"%v\", %v},\n", strings.ReplaceAll(k, "\"", "\\\""), v)
		fmt.Println(v, k)
	}
}
