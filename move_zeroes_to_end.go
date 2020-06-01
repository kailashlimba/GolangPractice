package main

import ("fmt")

func main(){
	
	array:=[]int{ 1, 2 ,3 ,8, 0, 78, 0, 5,0, 45}
	
	i:=0
	
	for _,value:= range array {
		if value != 0 {
			array[i] = value
			i++
		}
	}
	
	for ind:=i;ind<len(array);ind++ {
		array[ind] = 0
	}
	
	fmt.Println(array)
}
