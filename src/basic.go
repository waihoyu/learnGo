package main

import "fmt"

func variableZeroValue() {
	a := 1000
	s := "Learn Go"
	fmt.Printf("%d\n%s\n", a, s)
}

func variableInitialValue() {
	a := 3
	b := 4
	s := "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func main() {
	fmt.Println("Hello World")
	//variableZeroValue()
	//variableInitialValue()
	//variableTypeDeduction()
}
