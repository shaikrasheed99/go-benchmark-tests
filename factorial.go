package factorial

func factorialUsingIteration(number int) int {
	fact := 1

	for i := 1; i <= number; i++ {
		fact = fact * i
	}

	return fact
}
