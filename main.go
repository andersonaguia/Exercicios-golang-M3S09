package main

import (
	"fmt"
	"my_package/arraySum"
)

func main() {
	myArray := []int{1,2,3,4}

	var result int

	result = arraySum.ArraySum(myArray)

	fmt.Println(result)
	
}
