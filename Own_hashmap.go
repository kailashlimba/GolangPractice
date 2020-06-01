package main

import (
	"fmt"
)

type Pair struct {
	Key   string
	Value int
}

func main() {

	input_map := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4}
		
	
	var result_map [] Pair
	
	for key, value := range input_map {
		result_map = append(result_map,Pair{key,value})
	}
	
	fmt.Println(result_map)

}
