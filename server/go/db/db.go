package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/m/go/model"
	"github.com/m/go/utils"
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

func GetUserInfoDB(address string) (*model.User, error) {
	user := &model.User{}
	err := db.QueryRow("SELECT id, wallet_address, human_verified, signature FROM users WHERE wallet_address = $1", address).Scan(&user.ID, &user.WalletAddress, &user.HumanVerified, &user.Signature)
	return user, err
}

// func SetVerifiedDB(user model.User) {
// 	fmt.Println(user)
// 	_, err := db.Exec("UPDATE users SET human_verified = $1 WHERE wallet_address = $2", user.HumanVerified, strings.ToLower(user.WalletAddress))
// 	utils.HandleErr(err)
// }

func SaveUserDB(user model.User) {
	_, err := db.Exec("INSERT INTO users (wallet_address, human_verified, signature) VALUES ($1, $2, $3)", user.WalletAddress, user.HumanVerified, user.Signature)
	utils.HandleErr(err)
}

func CheckSignatureDB(address string, signature string) bool {
	var dbSignature string
	address = strings.ToLower(address) // 소문자로 변환
	err := db.QueryRow("SELECT signature FROM users WHERE wallet_address = $1", address).Scan(&dbSignature)
	if err != nil {
		utils.HandleErr(err)
		return false
	}
	fmt.Println(dbSignature)
	fmt.Println(signature)
	return dbSignature == signature
}

func SaveOrderDB(order model.OrderReq) (*model.Order, error) {
	newOrder := &model.Order{}

	query := `INSERT INTO orders (wallet_address, amount, price, is_buy, status, updated_at)
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, wallet_address, amount, price, is_buy, status, updated_at`
	err := db.QueryRow(query, strings.ToLower(order.WalletAddress), order.Amount, order.Price, order.IsBuy, "open", utils.GetTime()).Scan(
		&newOrder.ID, &newOrder.WalletAddress, &newOrder.Amount, &newOrder.Price, &newOrder.IsBuy, &newOrder.Status, &newOrder.UpdatedAt)

	if err != nil {
		utils.HandleErr(err)
		return nil, err
	}

	return newOrder, nil
}

func GetSellOrder4MatchDB(orderReq model.OrderReq) ([]model.Order, error) {
	query := `
	SELECT id, wallet_address, amount, price, is_buy, status, updated_at
	FROM orders
	WHERE price <= $2
		AND is_buy = false
		AND wallet_address != $1
		AND status = 'open'
	ORDER BY price ASC, updated_at DESC
	` // 최저가로 오름차순, 겹치는 가격이라면 최신 주문으로 정렬

	rows, err := db.Query(query, orderReq.WalletAddress, orderReq.Price)
	if err != nil {
		return nil, fmt.Errorf("error fetching matching sell orders: %w", err)
	}
	defer rows.Close()

	var orders []model.Order
	for rows.Next() {
		var order model.Order
		err := rows.Scan(&order.ID, &order.WalletAddress, &order.Amount, &order.Price, &order.IsBuy, &order.Status, &order.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning order: %w", err)
		}
		orders = append(orders, order)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading order rows: %w", err)
	}
	return orders, nil
}

func RecordTradeLog() {
	fmt.Println("will be recorded")
}
