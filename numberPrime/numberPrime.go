package numberPrime

import "my_package/labMath"

func CheckPrimeNumbers(arr []int) []int {
	primeNumbers := []int{}

	for i := 0; i < len(arr); i++ {
		isPrime := labMath.IsNumberPrime(arr[i])
		if isPrime {
			primeNumbers = append(primeNumbers, arr[i])
		}
	}

	return primeNumbers
}
