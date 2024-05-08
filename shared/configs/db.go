package configs

import (
	"fmt"
	"log"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	DB struct {
		Host     string
		User     string
		Password string
		Database string
		Port     int
		SSLMode  bool
		TimeZone string
	}
)

func GetDB() *gorm.DB {
	// Database connection string

	dsn := EnvConfigVars.DatabaseUrl

	// Open the connection to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting generic database object: %v", err)
	}
	// defer sqlDB.Close()
	// Ping the database to check for connection
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	// Your database operations go here
	fmt.Println("Successfully connected to the database")

	return db
}

func NewDBFromURL(databaseURL string) (*DB, error) {
	fields := strings.Fields(databaseURL)
	params := make(map[string]string)

	for _, field := range fields {
		parts := strings.SplitN(field, "=", 2)
		if len(parts) != 2 {
			continue
		}
		params[parts[0]] = parts[1]
	}

	port := 5432 // Default port
	if val, ok := params["port"]; ok {
		port = atoi(val)
	}

	sslMode := strings.EqualFold(params["sslmode"], "enable")

	return &DB{
		Host:     params["host"],
		User:     params["user"],
		Password: params["password"],
		Database: params["dbname"],
		Port:     port,
		SSLMode:  sslMode,
		TimeZone: params["timezone"],
	}, nil
}

func atoi(s string) int {
	n := 0
	for _, c := range s {
		n = n*10 + int(c-'0')
	}
	return n
}

func (c *DB) ConnectionString() string {
	ssl := "disable"
	if c.SSLMode {
		ssl = "enable"
	}

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s&timezone=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
		ssl,
		c.TimeZone,
	)
}
