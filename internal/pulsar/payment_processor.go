package pulsar

import (
    // "math/rand"
    "encoding/json"
    // "time"

    "github.com/dawitel/addispay-project/internal/domain"

    "github.com/apache/pulsar-client-go/pulsar"
)

// // mock the status change based on the output of the payment logic implemetation
// func getRandomStatus(statuses []string) string {
//         // Seed the random number generator for better randomness
//         rand.Seed(time.Now().UnixNano())

//         // Generate a random index within the range of the statuses slice
//         randomIndex := rand.Intn(len(statuses))

//         // Return the status at the randomly selected index
//         return statuses[randomIndex]
// }

// PaymentProcessingFunction is a function that takes order in and processes payments
func PaymentProcessorFunc(ctx pulsar.FunctionContext, input []byte) error {
    
    // // mock status changes
    // statuses := [...]string{"FAILED", "PENDING", "SUCCESS"}
    // status := getRandomStatus(statuses)

    var order domain.Order
    
    err := json.Unmarshal(input, &order)
    if err != nil {
        logger.Error("Error unmarshaling order: %v", err)
        return err
    }

    // Simulate payment processing
    orderResponse := domain.OrderResponse{
        OrderID: order.OrderID,
        Status:  "Success",  // TODO: move this field to status
    }

    output, err := json.Marshal(orderResponse)
    if err != nil {
        logger.Error("Error marshaling payment response: %v", err)
        return err
    }

    return ctx.Output("payment-results-topic", output)
}
