package lib

import (
    "time"
    "fmt"
)

const (
    LimitOrder = "LimitOrder"
    MarketOrder = "MarketOrder"
    Buy = "Buy"
    Sell = "Sell"
)

type Order struct {
    Id int
    // Time is ISO8601: YYYY-MM-DDThh:mm:ss.ms
    Time time.Time
    Quantity int
    Price float32
    OwnerId int
    Type string
    Kind string
}
 
func NewLimitOrder(typ string, price float32, quantity int, owner int) *Order {
    t := time.Now()
    return &Order{Time: t, Quantity: quantity, Price: price, OwnerId: owner, Type: typ, Kind: LimitOrder}
}
 
func NewMarketOrder(typ string, price float32, quantity int, owner int) *Order {
    t := time.Now()
    return &Order{Time: t, Quantity: quantity, Price: price, OwnerId: owner, Type: typ, Kind: MarketOrder}
}

// String implementation of an order
func (o Order) String() string {
    return fmt.Sprintf("(%s) %s - #%d - %d @ %12.2f on %v", o.Kind, o.Type, o.Id, o.Quantity, o.Price, o.Time)
}

// Convert between currencies
func (o Order) convert(to string, from string) float32 {
    return 0.0
}
