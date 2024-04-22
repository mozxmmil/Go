package main

import (
	"fmt"
	
)

// Day 4
func PrintingPattern(input int) {
	for i := 1; i <= input; i++ {
		for j := 1; j <= input; j++ {
			fmt.Print(string('A' + i - 1))
		}
		fmt.Println()
	}
}
 
func main(){
    var  input int
	fmt.Println("Enter the number of rows: ")
   fmt.Scanln(&input)
   PrintingPattern(input)
}