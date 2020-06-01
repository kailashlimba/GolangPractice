package main

import (
	"fmt"
)



func main() {

	var input_string = "India is my country India "
	
	mymap := make(map[string]int)
	
	str:=""
	
	for _,value :=range input_string {
		
		if value != ' ' {
			str = str + string(value)
		}else{
			if str[len(str)-1] == ',' || str[len(str)-1] == ':' || str[len(str)-1]== '"' || str[len(str)-1] == '\'' || str[len(str)-1] == '.'{
				str = str[0:len(str)-1]
			}
			
			if str[0] == ',' || str[0] == ':' || str[0]== '"' || str[0] == '\'' ||  str[0]== '.' {
				str = str[1:len(str)-1]
			}
			mymap[str]++
			
			str = ""
			
		}		
	}
	
	cnt:=0
	
	
	for _,value := range mymap {
		if value == 1{
			cnt++
		}
	}
	fmt.Println("Number of Unique words: ", cnt)
}
