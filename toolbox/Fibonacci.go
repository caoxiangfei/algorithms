package toolbox

import "log"

func FibRecurs(n int) (f int) {
	if n < 0 {
		log.Fatal("input must equal or large than 0")
	}
	if n <= 1 {
		return n
	} else {
		return FibRecurs(n-1) + FibRecurs(n-2)
	}
}

func FibList(n int) (f int) {
	F := []int{0, 1}

	if n > 1 {
		for i := 2; i < n+1; i++ {
			F = append(F, F[i-1]+F[i-2])
		}
	}
	return F[n]
}
