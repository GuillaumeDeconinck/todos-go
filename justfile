dev:
  reflex -r '\.go$' -s go run cmd/api.go

unit_tests:
  go clean -testcache && go test ./internal/... ./pkg/...

integration_tests:
  go clean -testcache && go test ./tests/integration/...

launch_db:
  docker run -d -p 5432:5432 --rm --name postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=todos postgres:alpine