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
	Signature     string `db:"signature" json:"signature"`
	IsNewUser     bool
	IsVerified    bool
}
