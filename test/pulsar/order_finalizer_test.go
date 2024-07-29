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

// test for the order finalizer func
func TestOrderFinalizationFunc(t *testing.T) {
    orderResponse := domain.OrderResponse{
        OrderID: "123",
        Status:  "Success",
    }

    responseJSON, err := json.Marshal(orderResponse)
    assert.NoError(t, err)

    err = pulsar.OrderFinalizationFunc(context.TODO(), responseJSON)
    assert.NoError(t, err)
}
