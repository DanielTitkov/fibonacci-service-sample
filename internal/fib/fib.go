package fib

import "math/big"

func N(n int) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}

	var fib2, fib1 = big.NewInt(0), big.NewInt(1)

	for i := 1; i < n; i++ {
		// TODO: add context to stop very long calculations
		fib2.Add(fib2, fib1)
		fib1, fib2 = fib2, fib1
	}

	return fib1
}
