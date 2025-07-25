# The Chamber of Keys

An in-memory key-value store featuring:

- String and list operations
- TTL (Time-To-Live) support
- Automatic persistence to disk

---

## Getting Started

Ensure Go v1.19+ is installed.

Navigate to the the_chamber_of_keys directory and run the following command to build the server:

```
go build -o server ./cmd/server
```

### Start the Service

Use the following command to launch the server:

```
./server
```

The Open API specification is located at `docs/openapi.yaml` 

You can explore it interactively using the [Swagger Editor](https://editor.swagger.io/).

---

## Testing and Benchmarking

To execute all the tests:

```
go test ./... -v
```

To run benchmark tests on the Chamber of Keys package:

```
go test -bench=. -benchmem ./pkg/chamber_of_keys
```

## Author

- **Nikita Rai**  
  Contact: [nikrc15@gmail.com](mailto:nikrc15@gmail.com)


