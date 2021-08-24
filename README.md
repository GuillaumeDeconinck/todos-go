# Todos API

## What is it ?

Just like the [Todos API implemented with NodeJS](https://github.com/GuillaumeDeconinck/todos-fastify), this API's intent is exactly the same, but instead is implemented with Golang.

**This repository is a rough WIP.**

## Requirements

This api requires a Postgres instance running at 5432. The `app.env` file lists all the details about it, and this file can be modified to match your environment.

The commands described here are implemented with the help of

- `Just`, a "simpler" makefile ([Github](https://github.com/casey/just))
- `Reflex`, a autoreloader utility ([Github](https://github.com/cespare/reflex))

## Developing

For running the app, simply run the following command:

```sh
just dev
```

Without `Just`, the command is as follows:

```sh
go run cmd/api.go
```

## Testing

For running unit tests without cache, use the following command:

```sh
just tests
```

Without `Just`, the command is as follows:

```sh
go clean -testcache && go test ./...
```

### Env vars

For the tests, if you need to specify the env vars, you can add a `app.env` file in the `tests` folder. Default values are specified in `internal/api/configuration/configuration.go`.

## Next steps

This API is far from being ready. This section will be soon completed.
