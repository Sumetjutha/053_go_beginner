package main

import "fmt"

// Condition 2
func main() {
	var number int
	fmt.Print("ป้อนตัวเลข = ")
	fmt.Scanf("%d", &number)
	fmt.Println("ตัวเลข = ", number)

	if number%2 == 0 {
		fmt.Println("เลขคู่")
	} else {
		fmt.Println("เลขคี่")
	}
}
