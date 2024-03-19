package main

import "fmt"

func main() {
	var number int
	fmt.Print("ใส่ค่า : ")
	fmt.Scanf("%d", &number)

	for i := 1; i <= number; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
	for i := 1; i <= number; i++ {
		for j := number - 1; j >= i; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
