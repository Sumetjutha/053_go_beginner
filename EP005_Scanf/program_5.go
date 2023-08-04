package main

import "fmt"

// scanf
func main() {
	// var name string
	// fmt.Print("ป้อนชื่อนักเรียน = ")
	// fmt.Scanf("%s", &name)
	// fmt.Println("สวัสดี ", name)

	// var score int
	// fmt.Print("ป้อนคะแนนนักเรียน = ")
	// fmt.Scanf("%d", &score)
	// fmt.Println("คะแนนสอบ + คะแนนช่วย = ", score+10)

	var score float64
	fmt.Print("ป้อนคะแนนนักเรียน = ")
	fmt.Scanf("%f", &score)
	fmt.Println("คะแนนสอบ + คะแนนช่วย = ", score+10.55)
}
