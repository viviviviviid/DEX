package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/m/go/model"
)

var (
	db *sql.DB
)

func InitDB() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(".env 파일을 찾을 수 없습니다.")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=require", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("DB is connected")
}

func GetUserInfo(address string) (*model.User, error) {
	user := &model.User{}
	err := db.QueryRow("SELECT wallet_address, human_verified FROM users WHERE wallet_address = $1", address).Scan(&user.WalletAddress, &user.HumanVerified)
	return user, err
}

func SetUser(address string) error {
	_, err := db.Exec("INSERT INTO users (wallet_address, human_verified) VALUES ($1, $2)", address, false)
	return err
}

func SetVerified(address string) error {
	_, err := db.Exec("UPDATE users SET human_verified = $1 WHERE wallet_address = $2", true, address)
	return err
}
