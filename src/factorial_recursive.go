package main

import "fmt"
import "math/big"
import "time"

func main() {

    //n := big.NewInt(40)
    start := time.Now()
    r := big.NewInt(500)
    fmt.Println(factorial(r))
    fmt.Println("time:", time.Since(start))
}

func factorial(n *big.Int) (result *big.Int) {
    //fmt.Println("n = ", n)
    b := big.NewInt(0)
    c := big.NewInt(1)

    if n.Cmp(b) == -1 {
        result = big.NewInt(1)
    }
    if n.Cmp(b) == 0 {
        result = big.NewInt(1)
    } else {
        // return n * factorial(n - 1);
        result = new(big.Int)
        result.Set(n)
        result.Mul(result, factorial(n.Sub(n, c)))
    }
    return
}