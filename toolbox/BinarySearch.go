package toolbox

func MyBinarySearch(A []int, low, high, key int) (index int) {
	if high < low {
		return low - 1
	}
	mid := low + (high-low)/2
	if key == A[mid] {
		return mid
	} else if key < A[mid] {
		return MyBinarySearch(A, low, mid-1, key)
	} else {
		return MyBinarySearch(A, mid+1, high, key)
	}
}

func BinarySearchIt(A []int, low, high, key int) (index int) {
	for low <= high {
		midd := low + (high-low)/2

		if key == A[midd] {
			return midd
		} else if key < A[midd] {
			high = midd - 1
		} else {
			low = midd + 1
		}
	}
	return low - 1
}
