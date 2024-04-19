package main

import "fmt"

func FindPrimeNum (num int){
	for i:=2; i<=num; i++ {
		if(num % i == 0){
			fmt.Println("not a prime number")
			break
		}else{
			fmt.Println("prime number")
			break
		}
	}
}

func main() {
	//day 1 to learn about the basic of go
	//1 to print any value you can you can use fmt.Println()
	fmt.Println("hey i am mozammil 😁")
	
	// 2 data types in go first leant sting and how to declare a variable
	// also use const
	var name1 string = "mozammil"
	var name2 = "mozammil"
	name3 := "mozammil"
	const name4 = "mozammil"
	const name5 = "mozammil"
	name6 := "mozammil"
	fmt.Println(name1,name2,name3,name4,name5,name6)
	// remember const use only for constant value 
	
    // make a find a prime num 
	
	FindPrimeNum(6)
	// 3 start print 

	var input int;
	fmt.Scanln(&input)
	for i:= 1; i<= input; i++{
		for j:=1 ; j<=input ; j++{
			fmt.Print("*")
		}
		fmt.Println()
	}

}
