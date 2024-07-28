package pulsar

import (
    "encoding/json"
    
    "github.com/dawitel/addispay-project/intenal/util"
    "github.com/dawitel/addispay-project/internal/domain"

    "github.com/apache/pulsar-client-go/pulsar"
)

var logger = util.GetLogger()

// OrderProcessingFunction 
func OrderProcessorFunc(ctx pulsar.FunctionContext, input []byte) error {
    var order domain.Order
    err := json.Unmarshal(input, &order)
    if err != nil {
        logger.Error("Error unmarshaling order: %v", err)
        return err
    }

    // TODO: Add order validation, enrichment
    logger.InfoGeneral("Processing order: %+v", order)

    output, err := json.Marshal(order)
    if err != nil {
        logger.Error("Error marshaling order: %v", err)
        return err
    }

    // publish the results from order processing to the topic "processed-orders-topic"
    return ctx.Output("processed-orders-topic", output)
}
