package main

import (
	"database/sql"
	"fmt"
	"github.com/EugeneNail/gotus/internal/database"
	_ "github.com/lib/pq"
	"os"
	"path"
	"time"
)

func main() {
	//migrateCommand := flag.NewFlagSet("migrate", flag.ExitOnError)
	if len(os.Args) == 1 {
		migrate()
		return
	}

	switch os.Args[1] {
	case "create":
		create()
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

func create() {
	if len(os.Args) < 3 {
		fmt.Println("Expected name of the migration")
		return
	}

	now := time.Now().Format("2006_02_01_150405")

	file, err := os.Create(path.Join(buildDirectory(), fmt.Sprintf("%s.%s.up.sql", now, os.Args[2])))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file, err = os.Create(path.Join(buildDirectory(), fmt.Sprintf("%s.%s.down.sql", now, os.Args[2])))
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

func buildDirectory() string {
	return path.Join(os.Getenv("APP_ROOT"), "deploy", "migrations")
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

	fmt.Println("The table 'migrations' does not exist.")
	fmt.Print("Create one? [Y/n]: ")

	var answer string
	if _, err := fmt.Scan(&answer); err != nil {
		panic(err)
	}

	if answer != "Y" {
		fmt.Println("Exited due to missing migration table")
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

	fmt.Println("The table 'migrations' has been created")
}
