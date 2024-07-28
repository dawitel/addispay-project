package main

import (
    "net"
    "github.com/dawitel/addispay-project/internal/grpc"
    "github.com/dawitel/addispay-project/protogen/golang/order"
    "github.com/dawitel/addispay-project/intenal/util"

    "github.com/apache/pulsar-client-go/pulsar"
    "google.golang.org/grpc"
)

var logger = util.GetLogger()

func main() {
    // instantiate a new Pulsar client
    pulsarClient, err := pulsar.NewClient(pulsar.ClientOptions{
        URL: "pulsar://localhost:6650",
    })

    if err != nil {
        logger.Error("Could not create Pulsar client: %v", err)
    }
    defer pulsarClient.Close()

    // create a producer that pushes messages to the orders-topic
    producer, err := pulsarClient.CreateProducer(pulsar.ProducerOptions{
        Topic: "orders-topic",
    })
    if err != nil {
        logger.Error("Could not create Pulsar producer: %v", err)
    }
    defer producer.Close()

    // TCP listener for the gRPC server
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        logger.Error("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    proto.RegisterOrderServiceServer(grpcServer, grpc.NewOrderServiceServer(producer))

    logger.Info("gRPC server started sreving on: localhos:50051", )
    if err := grpcServer.Serve(lis); err != nil {
        logger.Error("gRPC server failed to serve: %v", err)
    }
}
