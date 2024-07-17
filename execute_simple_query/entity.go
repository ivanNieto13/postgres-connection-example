package execute_simple_query

import "time"

type Product struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Price     float32   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
