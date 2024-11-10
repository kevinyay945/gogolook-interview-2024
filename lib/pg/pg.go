package pg

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Init is initial postgresql db
func Init(pgURL string) error {
	var err error
	db, err = gorm.Open(postgres.Open(pgURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("db connect err: %v", err)
		return err
	}
	_db, err := db.DB()
	if err != nil {
		log.Fatalf("set db connection pool failed: %v", err)
	}
	_db.SetMaxIdleConns(2)
	_db.SetMaxOpenConns(5)
	_db.SetConnMaxLifetime(time.Hour)

	return err
}

func GetPGURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_QUERY_PARAMS"),
	)
}
func GetDB() *gorm.DB {
	pgURL := GetPGURL()
	return GetDBByConnectingString(pgURL)
}

func GetDBByConnectingString(pgURL string) *gorm.DB {
	for {
		if db != nil {
			return db
		}
		if err := Init(pgURL); err != nil {
			time.Sleep(5 * time.Second)
		}
	}
}
