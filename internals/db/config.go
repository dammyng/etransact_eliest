package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Hosts    string
	Database string
	Username string
	Password string
	Port     string
}

func Config(db *DBConfig) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db.Username,
		db.Password,
		db.Hosts,
		db.Port,
		db.Database)
	return dsn
}

func InitConfig() DBConfig {
	var dBConfig DBConfig

	log.Println(os.Getenv("Environment"))
	if os.Getenv("Environment") != "production" {
		LoadEnv()
	}
	dBConfig.Database = os.Getenv("DBDatabase")
	dBConfig.Hosts = os.Getenv("DBHost")
	dBConfig.Password = os.Getenv("DBPassword")
	dBConfig.Port = os.Getenv("DBPort")
	dBConfig.Username = os.Getenv("DBUser")
	return dBConfig
}

func LoadEnv() {
	log.Println("env loading...")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}
}
