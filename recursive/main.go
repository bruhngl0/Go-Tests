package main

import (

	"fmt"
)


func RecBin(arr []int, low, high, target int) int{

	//base-case-for-rec
	if(low > high){
		return -1
	}

	mid := (low+high)/2;

	if arr[mid] == target {
		return mid
	}else if arr[mid] < target {
		return RecBin(arr, mid+1, high, target)
	}else {
		return RecBin(arr, low, mid-1, target)
	}

}


func main (){
arr := []int{1,2,3,4,5,6,7,8};

target := 6;
result := RecBin(arr, 0, len(arr)-1, target)
if result != -1 {
	fmt.Printf("index of target %d found at index %d", target, result)
}else {
	fmt.Printf("not found")  
}
}