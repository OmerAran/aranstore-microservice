package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type PostgresDb struct {
	DB *sqlx.DB
}

func GetDBConnection() *PostgresDb {
	db := CreateDBConnection()
	return &PostgresDb{
		DB: db,
	}
}

func CreateDBConnection() *sqlx.DB {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file", err)
		return nil
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sqlx.Connect("postgres", connectionString)

	if err != nil {
		log.Fatal(err.Error())
	}

	createTableSQL := `
	CREATE TABLE customers (
	    id SERIAL PRIMARY KEY,
	    name VARCHAR(255),
	    email VARCHAR(255) UNIQUE,
	    address VARCHAR(255),
		status int DEFAULT 1
	);
	`
	_, err = db.Exec(createTableSQL)

	return db
}
