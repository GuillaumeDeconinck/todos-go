# Todos API

## What is it ?

Just like the [Todos API implemented with NodeJS](https://github.com/GuillaumeDeconinck/todos-fastify), this API's intent is exactly the same, but instead is implemented with Golang.

**This repository is a rough WIP.**

## Requirements

This api requires a Postgres instance running at 5432. The `app.env` file lists all the details about it, and this file can be modified to match your environment. A simple Postgres instance can be ran with the following command:

```sh
just launch_db
# or
docker run -d -p 5432:5432 --rm --name postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=todos postgres:alpine
```

The commands described here are implemented with the help of

- `Just`, a "simpler" makefile ([Github](https://github.com/casey/just))
- `Reflex`, a autoreloader utility ([Github](https://github.com/cespare/reflex))

## Developing

For running the app, simply run the following command:

```sh
just dev
# or
reflex -r '\.go$' -s go run cmd/api.go
# or without autoreload
go run cmd/api.go
```

## Testing

For running unit tests without cache, use the following command:

```sh
just test
# or
go clean -testcache && go test ./...
```

> Note that `go clean -testcache` is not specifically needed, but it allows to run the tests from scratch

### Env vars

For the tests, if you need to specify the env vars, you can add a `app.env` file in the `tests` folder. Default values are specified in `internal/api/configuration/configuration.go`.

## Next steps

This API is far from being ready. This section will be soon completed.
