package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// InitDB initializes the database connection
func InitDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	// database, err := sql.Open("postgres", os.Genenv("DATABASE_URL"))
	// postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full
	//	_ "github.com/lib/pq"

	if err != nil {
		log.Fatal("Error opening database connection:", err)
		return nil, err
	}

	// Check if database connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
		return nil, err
	}
	log.Println("Database connected successfully")
	return db, nil
}
