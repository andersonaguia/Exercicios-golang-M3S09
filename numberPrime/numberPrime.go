package numberPrime

func isNumberPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	for i := 3; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func CheckPrimeNumbers(arr []int) []int {
	primeNumbers := []int{}

	for i := 0; i < len(arr); i++ {
		isPrime := isNumberPrime(arr[i])
		if isPrime {
			primeNumbers = append(primeNumbers, arr[i])
		}
	}

	return primeNumbers
}
