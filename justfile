dev:
  reflex -r '\.go$' -s go run cmd/api.go

test:
  go clean -testcache && go test ./...

launch_db:
  docker run -d -p 5432:5432 --rm --name postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=todos postgres:alpine