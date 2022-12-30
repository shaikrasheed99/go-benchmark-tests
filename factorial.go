package factorial

func factorialUsingIteration(number int) int {
	fact := 1

	for i := 1; i <= number; i++ {
		fact = fact * i
	}

	return fact
}

func factorialUsingRecursion(number int) int {
	if number <= 1 {
		return 1
	}

	return number * factorialUsingRecursion(number-1)
}
