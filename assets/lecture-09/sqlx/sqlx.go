package main

// MAIN START OMIT

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `CREATE TABLE person (
    first_name text,
    last_name text,
    email text
)`

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

func main() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	// MAIN END OMIT
	// SEED START OMIT
	db.MustExec(schema)

	tx := db.MustBegin()
	tx.Exec(`
		INSERT INTO person (first_name, last_name, email) 
		VALUES ($1, $2, $3)`,
		"Joe", "Bot", "jb@xd.net",
	)
	tx.NamedExec(`
		INSERT INTO person (first_name, last_name, email) 
		VALUES (:first_name, :last_name, :email)`,
		&Person{"Maria", "Letta", "maria.letta@gophers.com"},
	)
	tx.Commit()

	// SEED END OMIT
	// SELECT START OMIT

	var people []Person
	db.Select(&people, "SELECT * FROM person ORDER BY first_name ASC")

	var joe Person
	err = db.Get(&joe, "SELECT * FROM person WHERE first_name=$1", "Joe")

	var person Person
	rows, err := db.Queryx("SELECT * FROM person")
	for rows.Next() {
		err := rows.StructScan(&person)
		if err != nil {
			log.Fatal(err)
		}
	}

	// SELECT END OMIT
	// MAP START OMIT

	rows, err = db.NamedQuery(
		"SELECT * FROM person WHERE first_name=:first_name",
		joe,
	)

	rows, err = db.NamedQuery(
		"SELECT * FROM person WHERE first_name=:fn",
		map[string]any{"fn": "Bob"},
	)

	result, err := db.NamedExec(`
		INSERT INTO person (first_name,last_name,email) 
		VALUES (:first,:last,:email)`,
		map[string]any{
			"first": "Brad",
			"last":  "Firz",
			"email": "bf@google.com",
		},
	)

	// MAP END OMIT
	// BATCH-INSERT START OMIT

	personStructs := []Person{
		{FirstName: "Rob", LastName: "Pike", Email: "rob@google.com"},
		{FirstName: "Russ", LastName: "Cox", Email: "rsc@go.dev"},
	}
	result, err = db.NamedExec(`
		INSERT INTO person (first_name, last_name, email)
        VALUES (:first_name, :last_name, :email)`,
		personStructs,
	)

	personMap := []map[string]any{
		{"first_name": "Robert", "last_name": "Griesemer", "email": "rg@go.dev"},
		{"first_name": "Ken", "last_name": "Thompson", "email": "ken@barbie.com"},
	}
	result, err = db.NamedExec(`
		INSERT INTO person (first_name, last_name, email)
        VALUES (:first_name, :last_name, :email)`,
		personMap,
	)

	// BATCH-INSERT END OMIT

	_ = result
}
