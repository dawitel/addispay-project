package grpcServerPkg

import (
    "context"
    "encoding/json"

    "github.com/dawitel/addispay-project/internal/domain"
    pb "github.com/dawitel/addispay-project/internal/proto/proto"
    "github.com/dawitel/addispay-project/internal/util"

    "github.com/apache/pulsar-client-go/pulsar"
)


var logger = util.GetLogger()
 
// 
type OrderServiceServer struct {
    pb.UnimplementedOrderServiceServer
    pulsarProducer pulsar.Producer
}

// NewOrderServiceServer returns a new instace of the OrderServiceServer 
func NewOrderServiceServer(producer pulsar.Producer) *OrderServiceServer {
    return &OrderServiceServer{pulsarProducer: producer}
}

// SubmitOrder submits order to the order service
func (s *OrderServiceServer) SubmitOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
    order := domain.Order{
        OrderID:  req.OrderId,
        ItemName: req.ItemName,
        Quantity: req.Quantity,
        Price:    req.Price,
    }

    orderJSON, err := json.Marshal(order)
    if err != nil {
        logger.Error("Error marshaling order: %v", err)
        return nil, err
    }

    _, err = s.pulsarProducer.Send(ctx, &pulsar.ProducerMessage{
        Payload: orderJSON,
    })
    if err != nil {
        logger.Error("Error sending order message: %v", err)
        return nil, err
    }

    return &pb.OrderResponse{OrderId: req.OrderId, Status: "submitted"}, nil
}
