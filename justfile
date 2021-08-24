dev:
  reflex -r '\.go$' -s go run cmd/api.go

tests:
  go clean -testcache && go test ./...