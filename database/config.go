package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Connection string format: "user=username dbname=databasename sslmode=disable"
	connStr := "user=newuser dbname=NodeDB sslmode=disable password=pass port=5432"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Try to ping the database to check the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database!")

	rows, err := db.Query("SELECT name FROM table1")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {
		// var id int
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Name: %s\n", name)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

}
