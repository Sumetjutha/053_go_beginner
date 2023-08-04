package main

import "fmt"

// ตัวดำเนินการทางคณิตศาสตร์
func main() {
	// num_1 := 10
	// num_2 := 3
	num_1, num_2 := 10, 3

	fmt.Println("ผลบวก = ", num_1+num_2)
	fmt.Println("ผลลบ = ", num_1-num_2)
	fmt.Println("ผลคูณ = ", num_1*num_2)
	fmt.Println("ผลหาร = ", num_1/num_2)
	fmt.Println("ผลหารเหลือเศษ = ", num_1%num_2)
}
