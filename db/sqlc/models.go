// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"time"
)

type Session struct {
	ID           int64     `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type Token struct {
	ID              int64     `json:"id"`
	UID             int64     `json:"u_id"`
	TokenName       string    `json:"token_name"`
	Symbol          string    `json:"symbol"`
	Supply          int64     `json:"supply"`
	ContractAddress string    `json:"contract_address"`
	CreateTime      time.Time `json:"create_time"`
	UpdateTime      time.Time `json:"update_time"`
	DeleteTime      time.Time `json:"delete_time"`
	IsActive        bool      `json:"is_active"`
}

type Transaction struct {
	ID              int64  `json:"id"`
	TransactionType string `json:"transaction_type"`
	// from wallet address
	FromAddress string `json:"from_address"`
	// to wallet address, null when contract deploy
	ToAddress    string    `json:"to_address"`
	TransferData string    `json:"transfer_data"`
	HashValue    string    `json:"hash_value"`
	CreateTime   time.Time `json:"create_time"`
	UpdateTime   time.Time `json:"update_time"`
	DeleteTime   time.Time `json:"delete_time"`
	IsActive     bool      `json:"is_active"`
}

type User struct {
	ID                   int64     `json:"id"`
	Username             string    `json:"username"`
	HashedPassword       string    `json:"hashed_password"`
	FullName             string    `json:"full_name"`
	Email                string    `json:"email"`
	WalletPublicAddress  string    `json:"wallet_public_address"`
	WalletPrivateAddress string    `json:"wallet_private_address"`
	CreateTime           time.Time `json:"create_time"`
	UpdateTime           time.Time `json:"update_time"`
	DeleteTime           time.Time `json:"delete_time"`
	IsActive             bool      `json:"is_active"`
}
