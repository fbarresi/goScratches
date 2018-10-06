package main

import (
	"fmt"
	"math/big"
)

func main() {
	var f big.Int
	f.MulRange(1, 500)
	fmt.Println(&f)
}
