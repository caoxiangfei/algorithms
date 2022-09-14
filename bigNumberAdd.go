package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n string
	var m string

	fmt.Scan(&n)
	fmt.Scan(&m)

	var bigN big.Int
	var bigM big.Int
	bigN.SetString(n, 10)
	bigM.SetString(m, 10)

	bigN.Add(&bigN, &bigM)

	fmt.Println(bigN.String())
}
