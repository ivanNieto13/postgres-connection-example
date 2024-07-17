package execute_simple_query

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
)

// Execute runs a simple query on a Postgres database.
func Execute(ctx context.Context) {
	// get db connection
	conn := getConnection(ctx)

	// close connection on finish execution
	defer conn.Close(ctx)

	// get query from environment variable
	query := os.Getenv("QUERY")

	// execute query
	rows, err := conn.Query(ctx, query)

	// handle query error
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ejecutando la consulta: %v\n", err)
		os.Exit(1)
	}

	// close rows connection on finish execution
	defer rows.Close()

	// iterate returned rows
	for rows.Next() {
		var p Product

		// scan returned data into entity struct
		err = rows.Scan(&p.Id, &p.Name, &p.Price, &p.CreatedAt)

		// handle scan error
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error escaneando fila: %v\n", err)

			continue
		}

		// format entity to json
		b, _ := json.MarshalIndent(p, "", "   ")

		// print json object
		fmt.Print(string(b))
	}
}
