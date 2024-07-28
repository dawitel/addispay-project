package pulsar

import (
    "encoding/json"
    
    "github.com/dawitel/addispay-project/internal/domain"
    
    "github.com/apache/pulsar-client-go/pulsar"
)

func OrderFinalizationFunction(ctx pulsar.FunctionContext, input []byte) error {
    var orderResponse domain.OrderResponse
    
    err := json.Unmarshal(input, &orderResponse)
    if err != nil {
        logger.Error("Error unmarshaling order response: %v", err)
        return err
    }

    // Logger Error order status with the order ID
    logger.InfoGeneral("Order ID: %s, Status: %s", orderResponse.OrderID, orderResponse.Status)

    return nil
}
