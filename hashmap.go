package main

import ("fmt")

func main(){
	

	array1 := [6]int{0, 1, 2, 3, 4, 5}
	array2 := [6]int{2, 3, 5, 7, 11, 13}
	
	myMap := make(map[int]int)
	
	for index,value := range array1{
		myMap[value] = array2[index]
	}
	
	fmt.Println(myMap)
	
}
