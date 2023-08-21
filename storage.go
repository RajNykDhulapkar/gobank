package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountById(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=gobank_db password=password sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

// databae init
func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	query := `
	CREATE SEQUENCE IF NOT EXISTS account_number_seq START WITH 10000000 INCREMENT BY 1;

	CREATE TABLE IF NOT EXISTS accounts (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(255) NOT NULL,
		last_name VARCHAR(255),
		number BIGINT NOT NULL DEFAULT nextval('account_number_seq') UNIQUE,
		balance DECIMAL(10, 2) NOT NULL DEFAULT 0,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW()
	); 
	
	CREATE OR REPLACE FUNCTION update_updated_at()
	RETURNS TRIGGER AS $$
	BEGIN
		NEW.updated_at := NOW();
		RETURN NEW;
	END;
	$$ LANGUAGE plpgsql;

	CREATE TRIGGER trigger_update_updated_at
	BEFORE UPDATE ON accounts
	FOR EACH ROW
	EXECUTE FUNCTION update_updated_at();
	`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateAccount(account *Account) error {
	return nil
}

func (s *PostgresStore) UpdateAccount(accout *Account) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
}

func (s *) GetAccountById(id int) (*Account, error) {
	return nil, nil
}
