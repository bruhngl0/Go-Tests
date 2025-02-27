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
	fmt.Println("life is good" + "luck aint shit")
	fmt.Println("1+1=", 1+1)

	fmt.Println(true && false)
	fmt.Println(true || false)

	arr := []int{0, 1, 0, 3, 12}
	shiftZeros(arr)
	fmt.Println(arr)

}
