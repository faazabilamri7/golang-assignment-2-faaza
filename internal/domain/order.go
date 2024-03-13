package domain

import "time"

type Order struct {
    ID          uint
    OrderedAt   time.Time
    CustomerName string
    Items       []OrderItem
}

type OrderItem struct {
    ID          uint
    OrderID     uint
    ItemCode    string
    Description string
    Quantity    int
}
