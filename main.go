package main

import (
	_ "github.com/joho/godotenv/autoload"

	"postgres-connection-example/execute_simple_query"

	"context"
)

func main() {
	// declare context background var
	ctx := context.Background()

	// execute query
	execute_simple_query.Execute(ctx)
}
