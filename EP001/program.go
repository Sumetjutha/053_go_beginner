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
	name = "Martin Odegaard"  // เปลี่ยนแปลงข้อมูลได้

	/*
	 ตัวแปร constant คือค่าคงที่ที่ไม่สามารถเปลี่ยนแปลงได้
	*/
	// const name_3 string = "Pipupu"
	// name_3 = "Goky"

	fmt.Println("คะแนนสอบ => ", score)
	fmt.Println("My name is", name)
	fmt.Println("Age is => ", age, "year")
	fmt.Println("สอบผ่านมั๊ย => ", isPass)
	fmt.Println("ชื่อจริงๆ => ", name_2)
	// fmt.Println("ชื่อที่สาม => ", name_3)
	fmt.Println("Hello John Tik")
}
