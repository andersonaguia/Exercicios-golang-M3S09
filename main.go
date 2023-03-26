package main

import (
	"fmt"
	"strings"
	"my_package/arraySum"
	"my_package/fruitMap"
	"my_package/highestElement"
	"my_package/numberPrime"
	"my_package/stringConcat"
	"my_package/week"
	"my_package/labMath"
)

var teste string

func main() {
	myArray := []int{1, 2, 38, 4}

	var result int

	fmt.Println("\n*** Somando os elementos do array ***")

	result = arraySum.ArraySum(myArray)

	fmt.Println(result)

	fmt.Println("\n *** Verificando o elemento de maior valor contido no array ***")

	result = highestElement.HighestElement(myArray)

	fmt.Println(result)

	fmt.Println("\n *** Buscando palavras no dicionário de frutas ***")

	var stringResult string

	var fruit string

	fmt.Print("Digite a fruta que deseja procurar: ")

	fmt.Scan(&fruit)

	stringResult = fruitMap.FindFruit(strings.ToLower(fruit))

	fmt.Printf("%s: %s", fruit, stringResult)
	fmt.Println("")

	fmt.Println("\n *** Concatenando Strings ***")

	var string1 string

	var string2 string

	fmt.Print("Digite a primeira palavra: ")

	fmt.Scan(&string1)

	fmt.Print("Digite a segunda palavra: ")

	fmt.Scan(&string2)

	stringResult = stringConcat.StringConcat(string1, string2)

	fmt.Println(stringResult)

	fmt.Println("\n *** Verificando dias da semana ***")

	var day int

	fmt.Print("Digite o número referente ao dia da semana: ")

	fmt.Scan(&day)

	stringResult = week.WhatADay(day)

	fmt.Println(stringResult)

	fmt.Println("\n *** Encontrando números primos em um array ***")

	numbers := []int{}

	for i := 1; i <= 100; i++{
		numbers = append(numbers, i)
	}

	var primeNumbers []int

	primeNumbers = numberPrime.CheckPrimeNumbers(numbers)

	fmt.Print("Array original: ")
	fmt.Println(numbers)

	fmt.Print("Novo array apenas com numeros primos: ")
	fmt.Println(primeNumbers)

	fmt.Println("\n *** Pacote LabMath para verificar se o número é primo ***")

	var number int

	var isPrime bool

	fmt.Print("Digite o número desejado: ")
	fmt.Scan(&number)

	isPrime = labMath.IsNumberPrime(number)

	fmt.Printf("O numero %d é primo?: %t\n", number, isPrime)

}
