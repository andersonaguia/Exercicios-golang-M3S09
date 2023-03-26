package main

import (
	"fmt"
	"strings"
	"my_package/arraySum"
	"my_package/fruitMap"
	"my_package/highestElement"
	"my_package/stringConcat"
	"my_package/week"	
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

	fmt.Println("Concatenando Strings")

	var string1 string

	var string2 string

	fmt.Print("Digite a primeira palavra: ")

	fmt.Scan(&string1)

	fmt.Print("Digite a segunda palavra: ")

	fmt.Scan(&string2)

	stringResult = stringConcat.StringConcat(string1, string2)

	fmt.Println(stringResult)

	fmt.Println("Verificando dias da semana")

	var day int

	fmt.Print("Digite o número referente ao dia da semana: ")

	fmt.Scan(&day)

	stringResult = week.WhatADay(day)

	fmt.Println(stringResult)
}
