## Testing
### Unit tests
execute `go test -v ./users`

### Inegration tests
execute `docker compose -f docker-compose.integration.yml up api` to start databases and api
and run `go test -v ./intTest` to run actual tests

### Benchmark tests
execute `go test -bench=. ./pkg/hasher/.` to run tests for hasher package
