package main

import "fmt"

//Day 3 to printing pattern question
func PrintingPattern (input int){
	
	for i:= 1; i<=input; i++{
		for j:=1; j<=i; j++{
			print(i+j-1)
		}
		fmt.Println()
	}
}
func main() {
     var input int 
	 fmt.Println("enter Your num")
	 fmt.Scan(&input)
	 PrintingPattern(input)
}