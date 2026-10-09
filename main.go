package main

import (
	"fmt"

	"github.com/mstepan/leetcode/medium"
)

func main() {

	str := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"

	result := medium.FindRepeatedDnaSequences(str)

	if result == nil {
		fmt.Println("No results")
	} else {
		for _, single_seq := range result {
			fmt.Println(single_seq)
		}
	}

	fmt.Printf("Main done...")

}
