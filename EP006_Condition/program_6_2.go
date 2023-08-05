package main

import "fmt"

// Condition 2
func main() {
	var number int
	fmt.Print("ป้อนตัวเลข = ")
	fmt.Scanf("%d", &number)
	fmt.Println("ตัวเลข = ", number)

	if number == 1 {
		fmt.Println("เปิดบัญชีใหม่")
	} else if number == 2 {
		fmt.Println("ฝากเงิน - ถอนเงิน")
	} else {
		fmt.Println("ข้อมูลไม่ถูกต้อง")
	}
}
