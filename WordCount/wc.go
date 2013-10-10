package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	mm := make(map[string]int)
	elem := 1
	sfields := strings.Fields(s)
	for _, v := range sfields {
		elem = mm[v]
		elem++
		mm[v] = elem
	}

	return mm
}

func main() {
	wc := make(map[string]int)
	wc = WordCount("I am Learning go!")
	fmt.Println(wc)
	wc = WordCount("I I I test Word Count!")
	fmt.Println(wc)
}

