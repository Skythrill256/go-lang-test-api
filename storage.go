package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountById(int) (*Account, error)
}

type PostgressStore struct {
	db *sql.DB
}

func NewPostgressStore() (*PostgressStore, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgressStore{db: db}, nil
}

func (s *PostgressStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgressStore) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(50),
		last_name VARCHAR(50),
		number SERIAL,
		balance INT DEFAULT 0, 
		created_at TIMESTAMP
	)`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgressStore) CreateAccount(acc *Account) error {
	query := `INSERT INTO account
	(first_name, last_name, number, balance, created_at)
	VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.Exec(
		query,
		acc.FirstName,
		acc.LastName,
		acc.Number,
		acc.balance,
		acc.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgressStore) UpdateAccount(a *Account) error {
	return nil
}

func (s *PostgressStore) DeleteAccount(id int) error {
	return nil
}

func (s *PostgressStore) GetAccountById(id int) (*Account, error) {
	return nil, nil
}

func (s *PostgressStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("SELECT * FROM account")
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Ensure rows are closed after querying

	var accounts []*Account

	for rows.Next() {
		account := new(Account)
		err := rows.Scan(&account.ID, &account.FirstName, &account.LastName, &account.Number, &account.balance, &account.CreatedAt)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	// Check for errors from iterating over rows.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}
