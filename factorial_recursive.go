package main

import (
	"fmt"
	"math/big"
)

func main() {
	var x big.Int
	x.SetUint64(500)
	fmt.Println(Factorial(x))
}

func Factorial(n big.Int) (result big.Int) {
	var zero big.Int

	if n == zero {
		var x big.Int
		x.SetUint64(1)
		return x
	}

	result = n.Mul(Factorial(n - 1))
	return result

}
