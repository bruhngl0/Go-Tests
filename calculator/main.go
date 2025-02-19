package main
import (
	"fmt"
	"calculator/calculator"

)
func main(){
	result, err := calculator.Divide(2,0)
	if err != nil{
		fmt.Printf("Error: %s", err)
	}else{
		fmt.Printf("answer to your question is %f", result)
	}
		
	

}