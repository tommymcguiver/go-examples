package main

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"

	"fmt"
	"log"
	"reflect"

	"github.com/golang-migrate/migrate"
 "github.com/golang-migrate/migrate/v4/database/postgres"`


	mydb "github.com/tommymcguiver/go-examples/kyleconroy/sqlc/internal/db"
)

func main() {

	// connStr := "user=kenm host=localhost dbname=authors sslmode=disable"
	postgres, err := sql.Open("postgres", conectString)
	if err != nil {
		log.Fatal(err)
	}
	// Create driver instance from db.
	// Check each driver if it supports the WithInstance function.
	// `import "github.com/golang-migrate/migrate/v4/database/postgres"`
	instance, err := dStub.WithInstance(postgres, &dStub.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Read migrations from /home/mattes/migrations and connect to a local postgres database.
	m, err := migrate.NewWithInstance()
	if err != nil {
		log.Panic(err)
	}

	// Migrate all the way up ...
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	queries := mydb.New(postgres)
	ctx := context.Background()

	// list all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(authors)

	// create an author
	insertedAuthor, err := queries.CreateAuthor(ctx, mydb.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio:  sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(insertedAuthor)

	// get the author we just inserted
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		panic(err)
	}
	// prints true
	fmt.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))
}
