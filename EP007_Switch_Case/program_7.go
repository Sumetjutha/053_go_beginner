package main

import "fmt"

// Switch..Case
func main() {
	var number int
	fmt.Print("ป้อนตัวเลข = ")
	fmt.Scanf("%d", &number)
	fmt.Println("ตัวเลข = ", number)

	switch number {
	case 1:
		fmt.Println("เปิดบัญชีใหม่")
	case 2:
		fmt.Println("ฝากเงิน - ถอน")
	default:
		fmt.Println("ข้อมูลไม่ถูกต้อง")
	}
}
