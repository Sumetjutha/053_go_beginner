package main

import "fmt"

// แสดงผลชนิดของข้อมูล
func main() {
	name := "John Tik"
	age := 36
	score := 95.8
	isPass := true

	fmt.Printf("My name is %v \n", name)
	fmt.Printf("Type name = %T \n", name)

	fmt.Printf("Age = %v \n", age)
	fmt.Printf("Type age = %T \n", age)

	fmt.Printf("Score = %v \n", score)
	fmt.Printf("Type score = %T \n", score)

	fmt.Printf("Exam is = %v \n", isPass)
	fmt.Printf("Type isPass = %T \n", isPass)
}
