package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "admin"
	password = "password"
	dbname   = "messagepost"
)

type DBHandler struct {
	db *sql.DB
}

// NewDBHandler initializes the DBHandler and establishes a connection to the database
func NewDBHandler() (*DBHandler, error) {
	var connStr string = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		user, password, dbname, host, port)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	// Verify if the database connection works
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %v", err)
	}

	log.Println("Successfully connected to the database!")

	return &DBHandler{db: db}, nil
}

func (s *DBHandler) GetDB() *sql.DB {
	return s.db
}

// create table if not exist
func (s *DBHandler) CreateTables() error {
	query := `
    CREATE TABLE IF NOT EXISTS messages (
        id SERIAL PRIMARY KEY,
        body TEXT NOT NULL,
        created_at TIMESTAMPTZ DEFAULT NOW()
    );
    `
	_, err := s.db.Exec(query)
	if err != nil {
		return fmt.Errorf("unable to create tables: %v", err)
	}

	log.Println("tables created or already exist.")
	return nil
}

// close db connection
// might need to implement handling sigkill to signal running this function
func (s *DBHandler) CloseDB() error {
	err := s.db.Close()
	if err != nil {
		return fmt.Errorf("unable to close the database connection: %v", err)
	}
	log.Println("database connection closed.")
	return nil
}
