#!/bin/bash

# Replace these with your actual registry and image names
DOCKER_REGISTRY="dawitel"
IMAGE_NAME="addispay-project"
TAG="latest"

echo "Deploying Docker image..."

# Log in to Docker registry (if required)
# docker login $DOCKER_REGISTRY

# Tag the image
docker tag grpc-server:latest $DOCKER_REGISTRY/$IMAGE_NAME:$TAG

# Push the image to the registry
docker push $DOCKER_REGISTRY/$IMAGE_NAME:$TAG

echo "Deployment completed."
