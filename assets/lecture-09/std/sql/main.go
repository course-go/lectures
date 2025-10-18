package main

// MAIN START OMIT
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // registers a "postgres" driver
)

func main() {
	connStr := "postgres://username:password@hostname/database"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	rows, err := db.Query("SELECT name FROM users WHERE age = $1", 21)

	// MAIN END OMIT
	// EXEC START OMIT

	result, err := db.Exec(
		"INSERT INTO users (name, age) VALUES ($1, $2)",
		"Gopher",
		27,
	)
	affected, err := result.RowsAffected()

	// EXEC END OMIT
	// QUERYROW START OMIT

	name := "Bob"
	row := db.QueryRow("SELECT age FROM users WHERE name = $1", name)

	var age int64
	err = row.Scan(&age)

	// QUERYROW END OMIT
	// QUERY START OMIT

	rows, err = db.Query("SELECT name FROM users WHERE age = $1", age)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s is %d\n", name, age)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// QUERY END OMIT
	// TRANSACTION START OMIT

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	result, err = tx.Exec(`UPDATE orders SET status = ? WHERE id = ?`, "fullfilled", 2059)
	if err != nil {
		_ = tx.Rollback()
		log.Fatal(err)
	}

	result, err = tx.Exec(`UPDATE users SET status = ? WHERE id = ?`, "loyal", 27)
	if err != nil {
		_ = tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	// TRANSACTION END OMIT

	_ = affected
}
