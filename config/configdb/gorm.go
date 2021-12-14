package configdb

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	dsn := newPGConfig().String()

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	return db, nil
}

type pgConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

func newPGConfig() *pgConfig {
	dbConfig := pgConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
	return &dbConfig
}

func (dbConfig *pgConfig) String() string {
	mode := os.Getenv("MODE")
	dsn := ""
	if mode == "production" {
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=enable TimeZone=Asia/Jakarta",
			dbConfig.Host,
			dbConfig.User,
			dbConfig.Password,
			dbConfig.DBName,
			dbConfig.Port,
		)
	} else {
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			dbConfig.Host,
			dbConfig.User,
			dbConfig.Password,
			dbConfig.DBName,
			dbConfig.Port,
		)
	}

	return dsn
}
