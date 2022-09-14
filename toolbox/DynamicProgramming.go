package toolbox

func DPChange(money int, coins []int) (min int) {
	var MinNumCoins = []int{0}
	for i := 1; i < money+1; i++ {
		MinNumCoins = append(MinNumCoins, money+1)
	}

	for i := 1; i < money+1; i++ {
		for _, coin := range coins {
			if i >= coin {
				NumCoins := MinNumCoins[i-coin] + 1
				if NumCoins < MinNumCoins[i] {
					MinNumCoins[i] = NumCoins
				}
			}
		}
	}
	return MinNumCoins[money]
}
