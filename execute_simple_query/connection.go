package execute_simple_query

import (
	"github.com/jackc/pgx/v5"

	"context"
	"fmt"
	"os"
)

// getConnection returns a database connection.
func getConnection(ctx context.Context) *pgx.Conn {
	// performs db connection
	connection, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))

	// handle connection error
	if err != nil {
		_, err = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	return connection
}
