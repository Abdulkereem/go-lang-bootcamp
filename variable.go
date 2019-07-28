package main

import (
	"fmt"
	"strconv"
)

func main() {
	//var a, b, c = 3, 4, "Hello world"

	//var d int(
	//d = 42
	//fmt.Println(d)
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)

	// new_func()
	// point()
	// condition()
	// switch_case()
	//loop()
	//results := function(1,2)
	//fmt.Println(function(3,10)*200)
	//finite_loop()
	// show("Hello")
	array()

}

func new_func() {
	const me int = 10
	fmt.Println(me)
}

func point() {
	x := 20
	fmt.Println(x)
}

func condition() {
	var a int = 1
	robot := "Wash"

	if a == 1 {
		fmt.Println("here is your pomo boss!!")
	} else if robot == "Wash" {
		fmt.Println("Wash plate mode activated!!!")
	}
}

func switch_case() {
	var grade string = "B"
	var marks int = 80

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	}
	switch {
	case grade == "A":
		fmt.Println("Excellent")
	case grade == "B":
		fmt.Println("Good!!")
	default:
		fmt.Println("inalid")
	}
	fmt.Printf("your grade is ", grade)

}

func loop() {
	for true {
		fmt.Println("Thank you Miss Blessing!!!")
	}
}

func finite_loop() {
	var a int = 10

	for a < 20 {
		str := strconv.Itoa(a)
		fmt.Printf("Value of " + str + "\n")
		a++
		if a > 15 {
			fmt.Printf("break \n")
			break
		}
	}
}

func function(a int, b int) int {

	return a + b

}

func array() {
	//var names [7]string

	// names[0] = "Yemi"
	// names[1] = "Aliyu(Laku Master)"
	// names[2] = "Bazu(Laku Inventor)"
	// names[3] = "Ebuka"
	// names[4] = "Armel"
	// names[5] = "Saeeda"
	// names[6] = "Ahmed"

	// //fmt.Println(names)
	// fmt.Println(names[2])

	names :=[7]string{"Yemi","Aliyu(Laku Master)","Bazu(Laku Inventor)","Ebuka","Armel","Saeeda","Ahmed"}
	fmt.Println(names[1])
	//fmt.Println(names)

}
