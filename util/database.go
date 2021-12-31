package util

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go-lang-server/config"
	"log"
)

func setupDatabase(config *config.DBConfig) (*sql.DB, error) {

	db, err := sql.Open(config.DBDriver,
		fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
			config.PostgresUser,
			config.PostgresPassword,
			config.PostgresHost,
			config.PostgresPort,
			config.PostgresDB))

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected")

	return db, nil
}

func ConnectDB(cfg config.DBConfig) error {
	_, dbErr := setupDatabase(&cfg)
	if dbErr != nil {
		log.Fatalln("Error in connecting to DB")
		return dbErr
	}

	return nil
}
