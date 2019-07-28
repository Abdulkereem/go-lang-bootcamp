package main

import (
	"fmt"
)

func main() {

	fmt.Println("Make a choice!! ")
	fmt.Println("-------------------")
	fmt.Println("Addition (1)")
	fmt.Println("Subtraction (2)")
	fmt.Println("Multiplication (3)")
	fmt.Println("Division (4)")
	fmt.Println("-------------------")
	var input string
	fmt.Scanln(&input)
	if input == "1" {
		fmt.Println("Enter the values you want to add")
		var value1 int
		var value2 int
		fmt.Scanln(&value1)
		fmt.Scanln(&value2)
		fmt.Println(add(value1, value2))
	} else if input == "2" {
		fmt.Println("Enter the values you want to minus")
		var value1 int
		var value2 int
		fmt.Scanln(&value1)
		fmt.Scanln(&value2)
		fmt.Println(minus(value1, value2))
	} else if input == "3" {
		fmt.Println("Enter the values you want to mutiply")
		var value1 int
		var value2 int
		fmt.Scanln(&value1)
		fmt.Scanln(&value2)
		fmt.Println(multi(value1, value2))
	} else if input == "4" {
		fmt.Println("Enter the values you want to divide")
		var value1 int
		var value2 int
		fmt.Scanln(&value1)
		fmt.Scanln(&value2)
		fmt.Println(div(value1, value2))
	}
	//fmt.Println(input)
	//user:=fmt.Scanln(request_keyboard())
}

func add(a int, b int) int {
	return a + b

}

func multi(a int, b int) int {
	return a * b
}

func minus(a int, b int) int {
	return a - b
}

func div(a int, b int) int {
	return a / b
}

// func request_keyboard() string{
// 	reader := bufio.NewReader(os.Stdin)
// 	return reader.ReadString('\n')
// }
