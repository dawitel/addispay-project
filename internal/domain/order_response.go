package domain

// OrderResponse represents the response of an order.
type OrderResponse struct {
    OrderID string `json:"order_id"`
    Status  string `json:"status"`
}
