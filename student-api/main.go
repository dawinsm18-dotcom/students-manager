package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	// Read connection string
	connString := os.Getenv("PG_DSN")

	if connString == "" {
		log.Fatal("PG_DSN environment variable not set")
	}

	// Connect to PostgreSQL
	pool, err := pgxpool.New(
		context.Background(),
		connString,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer pool.Close()

	// Test connection
	err = pool.Ping(context.Background())

	if err != nil {
		log.Fatal("Connection failed:", err)
	}

	fmt.Println("PostgreSQL connected successfully!")

	// Check database name
	var dbName string

	err = pool.QueryRow(
		context.Background(),
		"SELECT current_database()",
	).Scan(&dbName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database:", dbName)

	// Retrieve student records
	rows, err := pool.Query(
		context.Background(),
		"SELECT id, name FROM example",
	)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	// Iterate through records
	for rows.Next() {

		var id int
		var name string

		err := rows.Scan(
			&id,
			&name,
		)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("ID:", id, " Name:", name)

	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

}
