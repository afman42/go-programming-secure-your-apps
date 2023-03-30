package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "afif123"
	dbPort   = "5432"
	dbName   = "postgres"
)

func LoadDB() *pgx.Conn {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)
	conn, err := pgx.Connect(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfull connected database")
	return conn
}
