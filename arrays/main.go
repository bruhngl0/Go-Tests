package main

import (
	"fmt"
)

func shiftZeros(nums []int) {
	lastNonZero := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[lastNonZero] = nums[lastNonZero], nums[i]
			lastNonZero++
		}

	}
}

func main() {

	fmt.Println("life is good" + "luck aint shit") //string-contacatenation
	fmt.Println("1+1=", 1+1)                       //inline operation
	fmt.Println(true && false)                     //false --output
	fmt.Println(true || false)                     //true --output
	fmt.Printf("%d - %b \n", 42, 42)               // %b for binary  %d for decimal base 10
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)   //%t for hexa

	var a string = "king"
	fmt.Println((a))

	var b int = 10
	fmt.Println(b)

	var c = 10
	fmt.Println(c)

	d := 10
	fmt.Println(d)

	e := "vari"
	fmt.Println(e)

	var n = 5000.00000
	var o = 22000.00 / n
	fmt.Println(o)

	fmt.Println(int64(o))

	const p = 5000.00000
	const q = 22000.00 / p
	fmt.Println(q)

	const g = "life is good"
	fmt.Println(g)

	arr := []int{0, 1, 0, 3, 12}
	shiftZeros(arr)
	fmt.Println(arr)

}
