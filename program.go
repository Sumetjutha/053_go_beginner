package main

import "fmt" //นำเอา Package เข้ามาทำงาน
/*
ให้ทำงานในฟังก์ชั่น main
มีการแสดงผลออกทางจอภาพ
โดยการแสดงผลจะอาศัยคำสัง fmtprintln
*/

// การนิยามตัวแปร

func main() {
	var name string = "John Tik"
	var age int = 36
	var score float32 = 95.8
	var isPass bool = true
	name_2 := "Sumet Jutajan" // เขียนแบบสั้นแบบนี้ได้

	fmt.Println("คะแนนสอบ => ", score)
	fmt.Println("My name is", name)
	fmt.Println("Age is => ", age, "year")
	fmt.Println("สอบผ่านมั๊ย => ", isPass)
	fmt.Println("ชื่อจริงๆ => ", name_2)
	fmt.Println("Hello John Tik")
}
