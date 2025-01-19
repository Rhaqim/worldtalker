# WorldTalker Translation Service

This project consists of a **gRPC-based Translation Service** with two components:  

1. **Server**: A Python gRPC server for handling translation requests.  
2. **Client**: A Go-based client that communicates with the server.

Both services run in Docker containers and are orchestrated using Docker Compose.

---

## Project Structure

```bash
.
├── client/                   # Go client source code
│   ├── service/              # Client application service folder
│   │   ├── translate.go      # Translation service
│   │   ├── websocket.go      # Websocket service
│   ├── main.go               # Entry point for the Go client
│   ├── go.mod                # Go module file
│   ├── go.sum                # Dependency lock file
│   ├── Dockerfile            # Dockerfile for the Go client
├── server/                   # Python server source code
│   ├── wtai/                 # Server application folder
│   │   ├── proto/            # gRPC protobuf definitions
│   │   │   ├── translate.proto
│   │   │   ├── translate_pb2.py
│   │   │   ├── translate_pb2_grpc.py
│   │   ├── server.py         # Entry point for the Python server
│   │   ├── __init__.py
│   ├── pyproject.toml        # Poetry dependency file
│   ├── poetry.lock           # Poetry lock file
│   ├── Dockerfile            # Dockerfile for the Python server
├── docker-compose.yml        # Docker Compose file
├── README.md                 # Project documentation
```

## Prerequisites

Ensure you have the following installed on your machine:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Getting Started

### Clone the repository

```bash
git clone https://github.com/Rhaqim/worldtalker.git
cd worldtalker
```

### Build and start the service

Use Docker Compose to build and start both the server and the client:

```bash
docker-compose up --build
```

This will:

- Start the Python server on port 50051.
- Start the Go client on port 8080.

## Configuration

### Environment Variables

You can configure the following environment variables in docker-compose.yml:

| Variable | Description | Default |
|----------|-------------|---------|
| `SERVER_HOST` | Hostname of the gRPC server | `server` |
| `SERVER_PORT` | Port number of the gRPC server | `50051` |

## Interactions

### Server

The Python server exposes a gRPC API defined in translate.proto. It listens on 0.0.0.0:50051.

### Client

The Go client communicates with the server using the translate.proto definition and sends translation requests.

### Translation Request

The client sends a translation request to the server with the following parameters:

- `message`: The text to be translated.
- `language_source`: The source language of the text.
- `language_target`: The target language to translate the text to.

### Translation Response

The server responds with the translated text.

## Testing

You can test the application on postman following the link below:

[![Run in Postman](https://run.pstmn.io/button.svg)](https://www.postman.com/warped-capsule-709231/workspace/worldtalker/collection/678d06983aee4976d23bbf8d?action=share&creator=17061476)
