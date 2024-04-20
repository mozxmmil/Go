package main

import "fmt"

func PrintStarONE (input int){
	// Day 2 again practice pattern question 
     for i:= 1; i<=input; i++{
		for j:=1; j<=input; j++{
			fmt.Print("*")
		}
		fmt.Println()
	 } 
}

// now second printPattern 
// 111
// 222
// 333
func PrintSecondPattern (input int){
	for i:= 1; i<=input; i++{
		for j:=1; j<=input; j++{
			fmt.Print(i)
		}
		fmt.Println()
	}
}
func main() {
	var input int
    fmt.Println("Enter Your Num")
	 fmt.Scanln(&input)

	// PrintStarONE(input)
	PrintSecondPattern(input)
}