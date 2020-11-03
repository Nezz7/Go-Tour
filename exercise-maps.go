package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	occ := make( map[string] int)
	words := strings.Fields(s)
	
	for i := 0; i < len(words); i++ {
		occ[words[i]]++
	}
	
	return  occ
}

func main() {
	occ := WordCount("I am learning Go!");
	fmt.Println(occ)
	fmt.Println("Expected output :")
	fmt.Println("map[string]int{\"Go!\":1, \"I\":1, \"am\":1, \"learning\":1}")
}
