package toolbox

func SelectionSort(A []int) (B []int) {
	for i := 0; i < len(A); i++ {
		minIndex := i
		for j := i + 1; j < len(A); j++ {
			if A[j] < A[minIndex] {
				minIndex = j
			}
		}
		A[i], A[minIndex] = A[minIndex], A[i]
	}
	return A
}

func MergeSort(A []int) (B []int) {
	if len(A) == 1 {
		return A
	}

	m := len(A) / 2

	C := A[0:m]
	D := A[m:]

	C = MergeSort(C)
	D = MergeSort(D)

	B = merge(C, D)
	return B
}

func merge(C, D []int) (E []int) {
	for len(C) > 0 && len(D) > 0 {
		c := C[0]
		d := D[0]
		if c <= d {
			E = append(E, c)
			C = append(C[:0], C[1:]...)
		} else {
			E = append(E, d)
			D = append(D[:0], D[1:]...)
		}
	}

	if len(C) > 0 {
		E = append(E, C...)
	}

	if len(D) > 0 {
		E = append(E, D...)
	}

	return E
}
