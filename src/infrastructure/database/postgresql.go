package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DatabaseConfig struct {
	Dialect  string
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type Database struct {
	*sqlx.DB
}

// NewPostgreSQLDBConnection return instance of DB Connection
func NewPostgreSQLDBConnection(c *DatabaseConfig) (database *Database, err error) {

	dbConn, err := sql.Open(c.Dialect, fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", c.Dialect, c.User, url.QueryEscape(c.Password), c.Host, c.Port, c.Database))
	if err != nil {
		log.Printf("Connect Database failed: %s", err.Error())
		return
	}

	db := sqlx.NewDb(dbConn, c.Dialect) // Create sqlx conn from sql.open

	// No need to ping manually to check if server is reachable
	// Sqlx call will it internally on connect
	if err = db.Ping(); err != nil {
		db.Close()
		log.Fatalf("PostgreSQL DB Server is unreachable. %s", err.Error())
		return
	}

	db.SetConnMaxLifetime(time.Minute * 15)
	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(30)

	database = &Database{db}

	log.Println("PostgreSQL Database connected")

	return
}
