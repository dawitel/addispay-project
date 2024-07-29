#!/bin/bash

PULSAR_URL="pulsar://localhost:6650"

echo "Deploying Pulsar functions..."

pulsar-admin functions create \
  --jar bin/orderProcessorFunc/order_processor.zip \
  --classname main \
  --inputs orders-topic \
  --output processed-orders-topic \
  --name order-processor-func \
  --pulsar-url $PULSAR_URL

pulsar-admin functions create \
  --jar bin/paymentProcessorFunc/payment_processor.zip \
  --classname main \
  --inputs processed-orders-topic \
  --output payment-results-topic \
  --name payment-processor-func \
  --pulsar-url $PULSAR_URL

pulsar-admin functions create \
  --jar bin/orderFinalyzerFunc/order_finalizer.zip \
  --classname main \
  --inputs payment-results-topic \
  --name order-finalizer-func \
  --pulsar-url $PULSAR_URL

echo "Pulsar functions deployed."
