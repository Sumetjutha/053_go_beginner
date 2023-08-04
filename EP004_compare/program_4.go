package main

import "fmt"

// ตัวดำเนินการเปรียบเทียบ
func main() {
	num_1, num_2 := 10, 3

	fmt.Println("เท่ากันหรือไม่ = ", num_1 == num_2)
	fmt.Println("ไม่เท่ากันจริงหรือไม่ = ", num_1 != num_2)
	fmt.Println(num_1, "มากกว่า", num_2, num_1 > num_2)
	fmt.Println(num_1, "น้อยกว่า", num_2, num_1 < num_2)
	fmt.Println(num_1, "มากกว่าหรือเท่ากับ", num_2, num_1 >= num_2)
	fmt.Println(num_1, "น้อยกว่าหรือเท่ากับ", num_2, num_1 <= num_2)
}
