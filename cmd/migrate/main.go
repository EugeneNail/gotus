package main

import (
	"database/sql"
	"fmt"
	"github.com/EugeneNail/gotus/internal/database"
	_ "github.com/lib/pq"
	"os"
)

const colorRed = "\033[31m"
const colorYellow = "\033[33m"
const colorGreen = "\033[32m"
const colorWhite = "\033[0m"

func main() {
	//command := flag.NewFlagSet("migrate", flag.ExitOnError)
	if len(os.Args) == 1 {
		migrate()
		return
	}

	switch os.Args[1] {
	case "create":
		fmt.Println("created")
	case "rollback":
		fmt.Println("rollbacked")
	default:
		fmt.Println("Expected 'create' or 'rollback' subcommands")
	}
}

func migrate() {
	db := database.Connect()
	createMigrationsTable(db)
}

func createMigrationsTable(db *sql.DB) {
	row := db.QueryRow(`
		SELECT COUNT(*) 
		FROM information_schema.tables
		WHERE table_name = 'migrations'
	`)

	var count int
	if err := row.Scan(&count); err != nil {
		panic(err)
	}

	if count != 0 {
		return
	}

	fmt.Println(colorize(colorYellow, "The table 'migrations' does not exist."))
	fmt.Print("Create one? [Y/n]: ")

	var answer string
	if _, err := fmt.Scan(&answer); err != nil {
		panic(err)
	}

	if answer != "Y" {
		fmt.Println(colorize(colorRed, "Exited due to missing migration table"))
		os.Exit(0)
	}

	_, err := db.Exec(`
		CREATE TABLE migrations
		(
			id          SERIAL PRIMARY KEY,
			name        TEXT NOT NULL,
			migrated_at TIMESTAMP
		)
	`)

	if err != nil {
		panic(err)
	}

	fmt.Println(colorize(colorGreen, "The table 'migrations' has been created"))
}

func colorize(color string, message string) string {
	return color + message + colorWhite
}
