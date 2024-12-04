# Thought Process

## Project Overview

This project is designed to handle incoming requests, track unique `id` values using Redis, and process them differently based on the extension type. There are three extensions:

1. **Extension 1**: Sends a `POST` request with the count of unique IDs to a specified endpoint.
2. **Extension 2**: Similar to Extension 1, but adds a `source` field (e.g., `load_balancer`).
3. **Extension 3**: Sends the count of unique IDs to a distributed streaming service.

## Implementation Approach

- **Configuration**: Environment variables are used for configuration (e.g., Redis address, log file path, endpoint).
- **Deduplication**: Redis ensures unique `id` values are counted by storing each ID temporarily with a TTL (1 minute).
- **Request Handling**: The main handler checks for the `id` parameter, ensures it's unique, and processes the request based on the extension type.
- **Extensions**:
  - **Extension 1 & 2**: Send the count of unique IDs to an endpoint. Extension 2 includes an additional `source` field.
  - **Extension 3**: Sends the unique ID count to a streaming service like Kafka or AWS Kinesis.
- **Metrics**: A simple logging mechanism writes the count to a log file (for Extensions 1 and 2) or to a streaming service (for Extension 3).

## Design Considerations

- **Modularity**: Each part of the system (configuration, Redis interaction, request processing) is modular, making it easy to extend.
- **Scalability**: Redis handles high-volume data efficiently, and streaming services in Extension 3 scale horizontally.
- **Concurrency**: The system is designed to handle multiple requests simultaneously using thread-safe data structures.

