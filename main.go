package main

import (
	"fmt"
	"my_package/arraySum"
	"my_package/highestElement"
)

var teste string

func main() {
	myArray := []int{1, 2, 38, 4}

	var result int

	fmt.Println("Somando os elementos do array")

	result = arraySum.ArraySum(myArray)

	fmt.Println(result)

	fmt.Println("Verificando o elemento de maior valor contido no array")

	result = highestElement.HighestElement(myArray)

	fmt.Println(result)



}
