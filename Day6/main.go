package main

import "fmt"

// Day 4
// ABC
// BCD
// CDE

// func PrintingPattern(input int){
// 	for i:=1 ; i<=input ; i++{
// 		for j:=1 ; j<=input; j++{
//          fmt.Print(string('A'+i+j-2))
// 		}
// 		fmt.Println()
// 	}
// }

//A
//BB
//CCC
//DDDD

func PrintingPattern (input int){
	for i:=1;i<=input;i++{
	  for	j:=1;j<=i;j++{
			fmt.Print(string('A'+i-1))
		}
		fmt.Println()
	}
}
func main() {
	var input int
	fmt.Println("Enter Your Num")
	fmt.Scanln(&input)
	 PrintingPattern(input)
}
