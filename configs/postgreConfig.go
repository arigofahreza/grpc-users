package configs

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PosgreConnection() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	postgreHost := os.Getenv("POSTGRE_HOST")
	postgrePort := os.Getenv("POSTGRE_PORT")
	postgreUser := os.Getenv("POSTGRE_USER")
	postgrePassword := os.Getenv("POSTGRE_PASSWORD")
	postgreDbname := os.Getenv("POSTGRE_DB")
	dsn := "host=" + postgreHost + " user=" + postgreUser + " password=" + postgrePassword + " dbname=" + postgreDbname + " port=" + postgrePort + " sslmode=disable TimeZone=Asia/Shanghai"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return conn, nil
}
