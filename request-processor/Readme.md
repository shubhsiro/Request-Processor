# Request Processor Project

This project processes incoming requests, deduplicates `id` values using Redis, and handles different extensions for request processing. It supports:

- Sending unique ID counts to a specified endpoint (Extensions 1 & 2).
- Sending unique ID counts to a distributed streaming service (Extension 3).

## Features

1. Deduplicates IDs using Redis with a 1-minute TTL.
2. Processes requests based on optional query parameters (`id`, `endpoint`, `extension`).
3. Logs unique counts to a file or streams the count to a service like Kafka.

---

## Prerequisites

1. **Go**: Install [Go](https://golang.org/dl/).
2. **Redis**: Install and run a Redis server. Ensure it's accessible at `localhost:6379` or configure the environment variable `REDIS_ADDR`.
3. **Kafka (Optional)**: For Extension 3, you need Kafka installed and running. Update `KAFKA_BROKER` in the environment variables with the broker's address.

---

## Installation and Setup

1. **Clone the Repository**:
    ```bash
    git clone <repository-url>
    cd request-processor
    ```

2. **Set Up Environment Variables**:
    Create a `.env` file (optional) or set environment variables directly:
    ```bash
    export SERVER_ADDR=:8080
    export LOG_FILE_PATH=./logs/app.log
    export REDIS_ADDR=localhost:6379
    export KAFKA_BROKER=localhost:9092 # Only for Extension 3
    export KAFKA_TOPIC=unique-id-count # Only for Extension 3
    ```

3. **Install Dependencies**:
    ```bash
    go mod tidy
    ```

4. **Run Redis**:
    Ensure the Redis server is running:
    ```bash
    redis-server
    ```

5. **Start the Server**:
    ```bash
    go run main.go
    ```

---

## Usage

### Endpoint

- **URL**: `/api/verve/accept`
- **Method**: `GET`
- **Query Parameters**:
    - `id` (required): A unique identifier for the request.
    - `endpoint` (optional): A URL to send the unique count to.
    - `extension` (optional): Determines the extension logic:
        - `1`: Logs the count to a file and sends it to the endpoint via POST.
        - `2`: Same as Extension 1, but adds a `source` field.
        - `3`: Sends the count to a streaming service like Kafka.

### Example Requests

1. **Basic Request**:
    ```bash
    curl "http://localhost:8080/api/verve/accept?id=123"
    ```

2. **With Extension 1**:
    ```bash
    curl "http://localhost:8080/api/verve/accept?id=123&endpoint=http://example.com&extension=1"
    ```
3. **With Extension 2**:
    ```bash
    curl "http://localhost:8080/api/verve/accept?id=123&endpoint=http://example.com&extension=2"
    ```

4. **With Extension 3 (Streaming)**:
    ```bash
    curl "http://localhost:8080/api/verve/accept?id=123&extension=3"
    ```

---

## Project Structure

- `main.go`: Entry point of the application.
- `config/`: Configuration handling.
- `handler/`: HTTP request handlers.
- `logger/`: Logging utilities.
- `metrics/`: Tracks unique requests and logs metrics.
- `storage/`: Redis client and deduplication logic.
- `request/`: Handles different extensions and request sending logic.

---

## Testing

- **Redis**: Verify Redis is storing keys:
    ```bash
    redis-cli keys *
    ```
- **Kafka**: Check Kafka topic messages using `kafka-console-consumer`:
    ```bash
    kafka-console-consumer --bootstrap-server localhost:9092 --topic unique-id-count --from-beginning
    ```

---

## Future Improvements

Implement comprehensive test cases for all modules.

---
