package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server *Server
	DB     *DB
}

type Server struct {
	Port string
}

type DB struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	TimeZone string
}

func GetConfig() *Config {
	loadEnv()

	db := setDB()
	server := setServer()

	return &Config{
		Server: server,
		DB:     db,
	}
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func checkEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("%s not set", key)
	}
	return value
}

func setServer() *Server {
	port := checkEnv("PORT")

	return &Server{
		Port: port,
	}
}

func setDB() *DB {
	host := checkEnv("DB_HOST")
	portDB := checkEnv("DB_PORT")
	user := checkEnv("DB_USER")
	password := checkEnv("DB_PASSWORD")
	dbName := checkEnv("DB_NAME")
	ssl := checkEnv("DB_SSLMODE")
	timeZone := checkEnv("DB_TIMEZONE")

	return &DB{
		Host:     host,
		Port:     portDB,
		User:     user,
		Password: password,
		DBName:   dbName,
		SSLMode:  ssl,
		TimeZone: timeZone,
	}
}
