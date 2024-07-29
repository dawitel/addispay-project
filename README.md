# Addispay Financial Services take-home project 
## Order Processing and Payment System

<p align="center">
  <a href="https://skillicons.dev">
    <img src="https://skillicons.dev/icons?i=git,docker,go,vim,aws" />
  </a>
</p>

This project implements an order processing and payment system using Apache Pulsar functions, gRPC, and Docker. The system is designed to handle high-concurrency order requests, process payments, and finalize orders, all in an asynchronous and scalable manner.

## Table of Contents

1. [Overview](#overview)
2. [Architecture](#architecture)
3. [Getting Started](#getting-started)
4. [Directory Structure](#directory-structure)
5. [Usage](#usage)
6. [Testing](#testing)
7. [Deployment](#deployment)
8. [License](#license)

## Overview

This system is composed of multiple services:
- **gRPC Server**: Receives order requests and forwards them to the order processing service.
- **Order Processing Service**: A Pulsar function that processes incoming orders.
- **Payment Processing Service**: A Pulsar function that processes payment for the orders.
- **Order Finalization Service**: A Pulsar function that finalizes the order processing based on payment results.

The services communicate via Apache Pulsar topics, with data being transferred in JSON format between Pulsar functions.

## Architecture

The architecture of this system involves several components working together:

- **gRPC Client**: Submits order requests.
- **gRPC Server**: Forwards order requests to the order processing system.
- **Apache Pulsar**: Facilitates messaging between the services.
- **Pulsar Functions**: Handle core business logic for order processing, payment processing, and order finalization.

### Architecture Diagram

![Architecture Diagram](docs/architecture.svg)
For more detail on the undelying architecture, thought processes, things that i learned, and components, see the [Architecture](docs/architecture.md) document.

## Getting Started

### Prerequisites

- **Docker** and **Docker Compose** installed on your machine.
- **Go** (1.19 or higher) for local development and testing.
- **Pulsar CLI** for managing Pulsar functions.

### Setup

1. **Clone the repository**:
    ```sh
    git clone https://github.com/dawitel/addispay-project.git
    cd order-payment-system
    ```

2. **Build the Docker images**:
    ```sh
    cd scripts && ./build.sh
    ```

3. **Start the services**:
    ```sh
    cd scripts && ./start.sh
    ```

4. **Deploy Pulsar functions**:
    ```sh
    cd scripts && ./deploy_pulsar_functions.sh
    ```

## Directory Structure

```plaintext
order-payment-system/
├── cmd/
│   ├── grpc_server/
│       └── main.go
├── configs/
│   ├── config.yml
├── docs/
│   ├── architecture.md
|   ├── architecture.svg
├── internal/
│   ├── domain/
│   │   ├── models.go
│   ├── grpc/
│   │   ├── order_services.go
│   ├── pulsar/
│   │   ├── order_processor.go
│   │   ├── payment_processor.go
│   │   ├── order_finalizer.go
│   ├── proto/
│   │   ├── order.proto
│   ├── util
│       ├── config.go
│       ├── logger.go
├── scripts/
│   ├── deploy.sh
|   ├── start.sh
|   ├── stop.sh
|   ├── test.sh
|   ├── build.sh
|   ├── deploy_pulsar_functions.sh
├── test/
│   ├── grpc/
│   │   ├── order_service_test.go
│   ├── pulsar/
│       ├── order_processor_test.go
│       ├── payment_processor_test.go
│       ├── order_finalizer_test.go
├── air.toml
├──.gitignore
├── .goreleaser.yml
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── LICENSE
├── Makefile
└── README.md
```

## Usage

1. **Submit Orders**: Use a gRPC client to submit order requests to the gRPC server. The server listens on port `50051`.

2. **Monitor Services**: Check the logs and status of services using `docker-compose logs` and `docker-compose ps`.

3. **Scaling**: Adjust the number of replicas or resources allocated to the services by modifying the `docker-compose.yml` file.

## Testing

To run tests, use the provided test script:

```sh
cd scripts && ./test.sh
```

This script runs unit tests for the gRPC server and Pulsar functions. Ensure that the environment is properly set up before running the tests.

## Deployment

1. **Build Docker Images**:
    ```sh
    cd scripts && ./build.sh
    ```

2. **Push to Docker Registry**: Use the `deploy.sh` script to push Docker images to your registry.
    ```sh
    cd scripts && ./deploy.sh
    ```

3. **Deploy**: Deploy the images to your target environment (e.g., Kubernetes, AWS ECS).



## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```

