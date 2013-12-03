package main

import (
    "fmt"
    "marketsim/lib"
)

func main() {
    lo := lib.NewLimitOrder(lib.Buy, 10.00, 10, 1)
    mo := lib.NewMarketOrder(lib.Sell, 12.00, 10, 1)
    fmt.Println(lo)
    fmt.Println(mo)
}
