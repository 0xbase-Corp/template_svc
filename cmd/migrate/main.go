package main

import (
	"embed"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	"github.com/0xbase-Corp/template_svc/shared/configs"
)

var (
	//go:embed migrations/*.sql
	src embed.FS
)

func main() {
	// Load environment variables from app.env
	configs.InitEnvConfigs()
	db := configs.GetDB()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("error getting SQL database:", err)
		return
	}

	// Initialize migration instance
	dir, err := iofs.New(src, "migrations")
	if err != nil {
		log.Fatal("error in migrations:", err)
		return
	}

	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		log.Fatal("error creating driver:", err)
		return
	}

	migrations, err := migrate.NewWithInstance("iofs", dir, "postgres", driver)
	if err != nil {
		log.Fatal("unable to fetch migrations:", err)
		return
	}

	if err := migrations.Up(); err != nil {
		log.Fatal("error in migrations:", err)
		return
	}

	log.Println("Migrations applied successfully.")
}
