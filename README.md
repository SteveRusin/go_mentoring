# Go Mentoring

## Getting Started

To start the project, follow these steps:

1. Run the command below:

    ```bash
    docker compose up
    ```

    This will start both databases and the API.

2. The service will be available on `localhost:3000`.

## Local developments
1. Run `make dev-http-service` - to start http api
1. Run `make dev-user-managment-service` - to start gRPC server
<!-- todo add example how to build rpc and update docker examples -->

## Example Requests

### Create a User

```bash
curl -X POST localhost:8080/user \
  -H "Content-Type: application/json" \
  -d '{"userName": "Steve", "password": "qwerty"}'
```

### Login

```bash
curl -X POST localhost:8080/user/login \
  -H "Content-Type: application/json" \
  -d '{"userName": "Steve", "password": "qwerty"}'
``` 

Make sure to replace `"userName"` and `"password"` with your desired values.

