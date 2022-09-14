package toolbox

func NaiveGCD(a, b int) (gcd int) {
	best := 0

	for d := 1; d < a+b; d++ {
		if a%d == 0 && b%d == 0 {
			best = d
		}
	}

	return best
}

//Let a′
//be the remainder when a is divided by
//b, then
//gcd(a, b) = gcd(a′, b) = gcd(b, a′).

func EuclidGCD(a, b int) (gcd int, err error) {
	if b == 0 {
		return a, err
	}

	a = a % b
	return EuclidGCD(b, a)
}
