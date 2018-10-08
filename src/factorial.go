package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	start := time.Now()
	var f big.Int
	f.MulRange(1, 1000)
	fmt.Println(&f)
	fmt.Println("time:", time.Since(start))
}
