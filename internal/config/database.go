package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
	dbOnce sync.Once

	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string
	POSTGRES_PORT     string
)

func init() {
	initDatabase()
}

func initDatabase() {
	SetupEnvFile()

	POSTGRES_USER = os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DB = os.Getenv("POSTGRES_DB")
	POSTGRES_PORT = os.Getenv("POSTGRES_PORT")

	dsn := "host=localhost user=" + POSTGRES_USER + " password=" + POSTGRES_PASSWORD + " dbname=" + POSTGRES_DB + " port=" + POSTGRES_PORT + " sslmode=disable TimeZone=UTC"
	dbOnce.Do(func() {
		var err error
		DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		MigrateDB(DBConn)
	})

}

func SetupEnvFile() {
	envFile := ".env"
	if _, err := os.Stat(envFile); os.IsNotExist(err) {
		envFile = "../.env"
	}
	if err := godotenv.Load(envFile); err != nil {
		panic("failed to load .env file")
	}
}
