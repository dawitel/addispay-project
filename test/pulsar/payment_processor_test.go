package pulsar_test

import (
    "context"
    "testing"
    "path/to/internal/domain"
    "path/to/internal/pulsar"
    "github.com/apache/pulsar-client-go/pulsar"
    "encoding/json"
    "github.com/stretchr/testify/assert"
)


// tests for the payment processor func
func TestPaymentProcessorFunc(t *testing.T) {
    order := domain.Order{
        OrderID:  "123",
        ItemName: "Book",
        Quantity: 2,
        Price:    20.0,
    }

    orderJSON, err := json.Marshal(order)
    assert.NoError(t, err)

    err = pulsar.PaymentProcessorFunc(context.TODO(), orderJSON)
    assert.NoError(t, err)
}
