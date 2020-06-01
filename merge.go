package main

import ("fmt")

func main(){
	

	array1 := []int{1, 2, 3, 4, 5}
	array2 := []int{2, 3, 5, 7, 11, 13}
	
	i:=0
	j:=0
	
	var result []int
	
	for i<len(array1) && j<len(array2) {
		if array1[i] < array2[j]{
			result = append(result, array1[i])
			i++
			
		}else{
			result = append(result, array2[j])
			j++
		}	
	}
	
	for i<len(array1) {
		result = append(result, array1[i])
		i++	
	}
	
	for j<len(array2) {
		result = append(result, array2[j])
		j++
		
	}
	
	fmt.Println(result)	
}
