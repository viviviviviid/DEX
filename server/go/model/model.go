package model

type User struct {
	ID            int    `db:"id" json:"id"`
	WalletAddress string `db:"wallet_address" json:"wallet_address"`
	HumanVerified bool   `db:"human_verified" json:"human_verified"`
}

type AuthReq struct {
	WalletAddress string `db:"wallet_address" json:"wallet_address"`
}

type AuthRes struct {
	WalletAddress string `db:"wallet_address" json:"wallet_address"`
	IsNewUser     bool
	IsVerified    bool
}
