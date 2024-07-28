#!/bin/bash

PULSAR_URL="pulsar://localhost:6650"

echo "Deploying Pulsar functions..."

pulsar-admin functions create \
  --jar path/to/order_processing_function.jar \
  --classname com.example.OrderProcessingFunction \
  --inputs orders-topic \
  --output processed-orders-topic \
  --name order-processing-function \
  --pulsar-url $PULSAR_URL

pulsar-admin functions create \
  --jar path/to/payment_processing_function.jar \
  --classname com.example.PaymentProcessingFunction \
  --inputs processed-orders-topic \
  --output payment-results-topic \
  --name payment-processing-function \
  --pulsar-url $PULSAR_URL

pulsar-admin functions create \
  --jar path/to/order_finalization_function.jar \
  --classname com.example.OrderFinalizationFunction \
  --inputs payment-results-topic \
  --name order-finalization-function \
  --pulsar-url $PULSAR_URL

echo "Pulsar functions deployed."
