package summultiples

func SumMultiples(limit int, divisors ...int) int {
	var sum int
	multiples := make([]bool, limit)
	for i := 1; i < limit; i++ {
		for _, d := range divisors {
			if !multiples[i] && d != 0 && i%d == 0 {
				multiples[i] = true
				sum += i
			}
		}
	}
	return sum
}
