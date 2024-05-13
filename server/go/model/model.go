package model

type User struct {
	ID            int    `db:"id" json:"id"`
	WalletAddress string `db:"wallet_address" json:"wallet_address"`
	HumanVerified bool   `db:"human_verified" json:"human_verified"`
	Signature     string `db:"signature" json:"signature"`
}

type AuthReq struct {
	WalletAddress string `db:"wallet_address" json:"wallet_address"`
	Signature     string `db:"signature" json:"signature"`
}

type AuthRes struct {
	WalletAddress string `db:"wallet_address" json:"wallet_address"`
	IsNewUser     bool
	IsVerified    bool
	IsSignature   bool
}

type OrderReq struct {
	WalletAddress string  `db:"wallet_address" json:"wallet_address"`
	TokenAddress  string  `db:"token_address" json:"token_address"`
	Amount        int     `db:"amount" json:"amount"`
	Price         float64 `db:"price" json:"price"`
	IsBuy         bool    `db:"is_buy" json:"is_buy"`
	Signature     string  `db:"signature" json:"signature"`
}

type Order struct {
	ID            int     `db:"id" json:"id"`
	WalletAddress string  `db:"wallet_address" json:"wallet_address"`
	Amount        int     `db:"amount" json:"amount"`
	Price         float64 `db:"price" json:"price"`
	IsBuy         bool    `db:"is_buy" json:"is_buy"`
	Status        string  `db:"status" json:"status"`
	UpdatedAt     string  `db:"updated_at" json:"updated_at"`
}
