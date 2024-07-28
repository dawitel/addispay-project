package grpc_test

import (
    "context"
    "testing"
    
    "github.com/dawitel/addispay-project/internal/grpc"
    "github.com/dawitel/addispay-project/proto"

    "github.com/apache/pulsar-client-go/pulsar"
    "github.com/stretchr/testify/assert"
    "google.golang.org/grpc"
)

func TestSubmitOrder(t *testing.T) {
    // Setup Pulsar mock client and producer
    client, err := pulsar.NewClient(pulsar.ClientOptions{
        URL: "pulsar://localhost:6650",
    })
    assert.NoError(t, err)
    defer client.Close()

    producer, err := client.CreateProducer(pulsar.ProducerOptions{
        Topic: "orders-topic",
    })
    assert.NoError(t, err)
    defer producer.Close()

    orderService := grpc.NewOrderServiceServer(producer)
    req := &proto.OrderRequest{
        OrderId:  "123",
        ItemName: "Book",
        Quantity: 2,
        Price:    20.0,
    }

    res, err := orderService.SubmitOrder(context.Background(), req)
    assert.NoError(t, err)
    assert.Equal(t, "123", res.OrderId)
    assert.Equal(t, "submitted", res.Status)
}
