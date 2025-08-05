package services

import (
	"BANK_TRANSACTION_APPLICATION/db"
	"BANK_TRANSACTION_APPLICATION/model"
	"errors"
)

func InsertAccount(acc model.Account) error {
	_, err := db.Conn.Exec(
		`INSERT INTO accounts (account_id, balance) VALUES ($1, $2)`,
		acc.AccountID, acc.Balance,
	)
	return err
}

func FetchAccount(id int64) (model.Account, error) {
	var acc model.Account
	err := db.Conn.QueryRow(
		`SELECT account_id, balance FROM accounts WHERE account_id = $1`,
		id,
	).Scan(&acc.AccountID, &acc.Balance)
	return acc, err
}

func ProcessTransaction(tx model.Transaction) error {
	conn := db.Conn
	txSQL, err := conn.Begin()
	if err != nil {
		return err
	}

	var sourceBal float64
	err = txSQL.QueryRow(
		`SELECT balance FROM accounts WHERE account_id = $1 FOR UPDATE`,
		tx.SourceAccountID,
	).Scan(&sourceBal)
	if err != nil {
		txSQL.Rollback()
		return err
	}

	if sourceBal < tx.Amount {
		txSQL.Rollback()
		return errors.New("insufficient funds")
	}

	_, err = txSQL.Exec(
		`UPDATE accounts SET balance = balance - $1 WHERE account_id = $2`,
		tx.Amount, tx.SourceAccountID,
	)
	if err != nil {
		txSQL.Rollback()
		return err
	}

	_, err = txSQL.Exec(
		`UPDATE accounts SET balance = balance + $1 WHERE account_id = $2`,
		tx.Amount, tx.DestinationAccountID,
	)
	if err != nil {
		txSQL.Rollback()
		return err
	}

	_, err = txSQL.Exec(
		`INSERT INTO transactions (source_account_id, destination_account_id, amount) VALUES ($1, $2, $3)`,
		tx.SourceAccountID, tx.DestinationAccountID, tx.Amount,
	)
	if err != nil {
		txSQL.Rollback()
		return err
	}

	return txSQL.Commit()
}
