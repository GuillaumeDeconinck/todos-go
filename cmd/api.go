package main

import (
	"github.com/GuillaumeDeconinck/todos-go/internal/api"
)

func main() {
	r := api.SetupApi()

	r.Run() // listen and serve on 0.0.0.0:8080
}
