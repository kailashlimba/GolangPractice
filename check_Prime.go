package main

import (
	"fmt"
)

func isPrime(number int, i int)bool{
	if number <= 2 {
		if number == 2 {
			return true
		}
		
		return false
	}
	
	if number%i == 0 {
		return false
	}
	
	if i*i > number {return true}
	
	return isPrime(number, i+1)
	
}	

func main() {
	var number = 17
	
	if isPrime(number,2) == true {
		fmt.Println("Number is prime")
	}else {
		fmt.Println("Not Prime")
	}
}
