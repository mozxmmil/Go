package main

import (
	
	"fmt"
)
func PrintPattern (input int){
	newnum:=0
	for i:=1; i<=input; i++{
		for j:=1; j<=input; j++{
			fmt.Print(string('A' +newnum))
			newnum++
		}
		fmt.Println()
	}
}

// main is the entry point of the program.
func main(){
	var input int
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&input)
	PrintPattern(input)
}