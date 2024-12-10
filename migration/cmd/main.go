package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	direction := flag.String("direction", "up", "down")
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

	m, err := migrate.New(
		"file://./migration",
		dsn,
	)
	if err != nil {
		log.Fatalf("failed to initialize migration: %v", err)
	}

	if *direction == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("failed to apply migration: %v", err)
		}
		log.Println("Migration applied successfully")
	} else if *direction == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("failed to roll back migration: %v", err)
		}
		log.Println("Migration rolled back successfully")
	} else {
		log.Fatalf("invalid direction: %v", *direction)
	}
}
