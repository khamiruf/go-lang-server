package util

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go-lang-server/config"
	"log"
)

type Database struct {
}

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

func loadConfig(path string) (*config.DBConfig, error) {
	config := new(config.DBConfig)

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)

	return config, nil
}

func ConnectDB() error {
	config, err := loadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
		return err
	}

	_, dbErr := setupDatabase(config)
	if dbErr != nil {
		log.Fatalln("Error in connecting to DB")
		return dbErr
	}

	return nil
}
