# hello-world
This is created for an assignment

## How to Run

### Build the project and
```sh
go build -o ./exec ./cmd/server
```

### Then run the binary
```sh
./exec
```

### Or else run the project directly from source using
```sh
go run ./cmd/server
```

### Example Request

You can test the running server with:

```sh
curl http://localhost:8080/hello-world?name=alice
```

Expected response:

```json
{"message":"Hello alice"}
```

### Run tests
```sh
go test -v -cover ./...
```
#### You can see the test results along with the coverages

## Assumptions

- Names may contain numbers in the middle; this is not treated as an error.
- Leading and trailing white spaces in the name are trimmed before processing.
