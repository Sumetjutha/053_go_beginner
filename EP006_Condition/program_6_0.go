package main

import "fmt"

// Condition
func main() {
	var score int
	fmt.Print("ป้อนคะแนนนักเรียน = ")
	fmt.Scanf("%d", &score)
	fmt.Println("คะแนนสอบ = ", score)

	// ประมวลผลคะแนน
	if score >= 50 {
		fmt.Println("สอบผ่าน")
	} else {
		fmt.Println("สอบไม่ผ่าน")
	}
}
