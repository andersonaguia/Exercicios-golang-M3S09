package main

import (
	"fmt"
	"strings"
	"my_package/arraySum"
	"my_package/fruitMap"
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

	var stringResult string

	var fruit string

	fmt.Print("Digite a fruta que deseja procurar: ")

	fmt.Scan(&fruit)

	stringResult = fruitMap.FindFruit(strings.ToLower(fruit))

	fmt.Printf("%s: %s", fruit, stringResult)
	fmt.Println("")

}
