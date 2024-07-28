package domain

// Order represents an order in the system.
type Order struct {
    OrderID  string  `json:"order_id"`
    ItemName string  `json:"item_name"`
    Quantity int32   `json:"quantity"`
    Price    float64 `json:"price"`
}
